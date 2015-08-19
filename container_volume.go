package libyaml

type ContainerVolume struct {
	HostPath      string `yaml:"host_path" json:"host_path"`
	ContainerPath string `yaml:"container_path" json:"container_path"`
	Permission    string `yaml:"permission" json:"permission"`
	Owner         string `yaml:"owner" json:"owner"`
}
