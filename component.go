package libyaml

const (
	ClusterStrategySpread = "spread"
	ClusterStrategyPack   = "pack"
)

type Component struct {
	Name             string                    `yaml:"name" json:"name"`
	Tags             []string                  `yaml:"tags" json:"tags"`
	Conflicts        []string                  `yaml:"conflicts" json:"conflicts"`
	Cluster          bool                      `yaml:"cluster" json:"cluster"`
	ClusterHostCount ComponentClusterHostCount `yaml:"cluster_host_count" json:"cluster_host_count"`
	Containers       []*Container              `yaml:"containers" json:"containers"`
}

type ComponentClusterHostCount struct {
	Min               minInt1 `yaml:"min" json:"min"`
	Max               uint    `yaml:"max,omitempty" json:"max"` // 0 == unlimited
	ThresholdHealthy  uint    `yaml:"threshold_healthy" json:"threshold_healthy"`
	ThresholdDegraded uint    `yaml:"threshold_degraded,omitempty" json:"threshold_degraded"` // 0 == no degraded state
}

func (c *Component) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var m marshallerComponent
	if err := unmarshal(&m); err != nil {
		return err
	}
	m.decode(c)

	if c.Cluster {
		if c.ClusterHostCount.Min == 0 {
			c.ClusterHostCount.Min = 1
		}
		if c.ClusterHostCount.ThresholdHealthy == 0 {
			c.ClusterHostCount.ThresholdHealthy = 1
		}
	}

	return nil
}

func (c Component) MarshalYAML() (interface{}, error) {
	if !c.Cluster {
		m := nonclusterableComponent{}
		m.encode(c)
		return m, nil
	}

	m := marshallerComponent{}
	m.encode(c)
	return m, nil
}

type marshallerComponent Component

func (m *marshallerComponent) encode(c Component) {
	m.Name = c.Name
	m.Tags = c.Tags
	m.Conflicts = c.Conflicts
	m.Cluster = c.Cluster
	m.ClusterHostCount = c.ClusterHostCount
	m.Containers = c.Containers
}

func (m marshallerComponent) decode(c *Component) {
	c.Name = m.Name
	c.Tags = m.Tags
	c.Conflicts = m.Conflicts
	c.Cluster = m.Cluster
	c.ClusterHostCount = m.ClusterHostCount
	c.Containers = m.Containers
}

type nonclusterableComponent struct {
	Name       string       `yaml:"name" json:"name"`
	Tags       []string     `yaml:"tags" json:"tags"`
	Conflicts  []string     `yaml:"conflicts" json:"conflicts"`
	Cluster    bool         `yaml:"cluster" json:"cluster"`
	Containers []*Container `yaml:"containers" json:"containers"`
}

func (m *nonclusterableComponent) encode(c Component) {
	m.Name = c.Name
	m.Tags = c.Tags
	m.Cluster = false
	m.Conflicts = c.Conflicts
	m.Containers = c.Containers
}
