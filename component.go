package libyaml

type Component struct {
	Name             string                    `yaml:"name" json:"name"`
	Tags             []string                  `yaml:"tags" json:"tags"`
	Conflicts        []string                  `yaml:"conflicts" json:"conflicts"`
	Cluster          BoolString                `yaml:"cluster" json:"cluster"`
	ClusterHostCount ComponentClusterHostCount `yaml:"cluster_host_count" json:"cluster_host_count"`
	HostRequirements ComponentHostRequirements `yaml:"host_requirements" json:"host_requirements"`
	LogOptions       LogOptions                `yaml:"logs" json:"logs"`
	HostVolumes      []*HostVolume             `yaml:"host_volumes" json:"host_volumes"`
	Containers       []*Container              `yaml:"containers" json:"containers" validate:"dive,exists"`
}

type ComponentClusterHostCount struct {
	// Strategy = "autoscale" api version >= 2.7.0
	// Strategy = "random" api version >= 2.5.0
	Strategy          string `yaml:"strategy,omitempty" json:"strategy,omitempty" validate:"omitempty,clusterstrategy"`
	Min               uint   `yaml:"min" json:"min"`
	Max               uint   `yaml:"max,omitempty" json:"max"` // 0 == unlimited
	ThresholdHealthy  uint   `yaml:"threshold_healthy" json:"threshold_healthy"`
	ThresholdDegraded uint   `yaml:"threshold_degraded,omitempty" json:"threshold_degraded"` // 0 == no degraded state
}
