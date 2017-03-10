package libyaml

type HostRequirements struct {
	DockerVersion     string `yaml:"docker_version" json:"docker_version" validate:"omitempty,dockerversion"`
	ReplicatedVersion string `yaml:"replicated_version" json:"version" validate:"omitempty,semverrange"`
	CPUCores          uint   `yaml:"cpu_cores" json:"cpu_cores"`
	CPUMhz            uint   `yaml:"cpu_mhz" json:"cpu_mhz"`
	Memory            string `yaml:"memory" json:"memory" validate:"bytes"`
	DiskSpace         string `yaml:"disk_space" json:"disk_space" validate:"bytes"`
}

type ComponentHostRequirements struct {
	CPUCores  uint   `yaml:"cpu_cores" json:"cpu_cores"`
	CPUMhz    uint   `yaml:"cpu_mhz" json:"cpu_mhz"`
	Memory    string `yaml:"memory" json:"memory" validate:"bytes"`
	DiskSpace string `yaml:"disk_space" json:"disk_space" validate:"bytes"`
}
