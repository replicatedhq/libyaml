package libyaml

type ContainerVolume struct {
	HostPath             string `yaml:"host_path" json:"host_path" validate:"required,absolutepath"`
	ContainerPath        string `yaml:"container_path" json:"container_path" validate:"required,absolutepath"`
	Permission           string `yaml:"permission" json:"permission"`
	Owner                string `yaml:"owner" json:"owner"`
	IsExcludedFromBackup string `yaml:"is_excluded_from_backup" json:"is_excluded_from_backup"`
}
