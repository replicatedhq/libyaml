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
	HostRequirements ComponentHostRequirements `yaml:"host_requirements" json:"host_requirements"`
	HostVolumes      []*HostVolume             `yaml:"host_volumes" json:"host_volumes"`
	Containers       []*Container              `yaml:"containers" json:"containers" validate:"dive,exists"`
}

type ComponentClusterHostCount struct {
	Min               uint `yaml:"min" json:"min"`
	Max               uint `yaml:"max,omitempty" json:"max"` // 0 == unlimited
	ThresholdHealthy  uint `yaml:"threshold_healthy" json:"threshold_healthy"`
	ThresholdDegraded uint `yaml:"threshold_degraded,omitempty" json:"threshold_degraded"` // 0 == no degraded state
}
