package libyaml

type ContainerExtraHost struct {
	Hostname string `yaml:"name" json:"name" validate:"required"`
	Address  string `yaml:"address" json:"address" validate:"required"`
	When     string `yaml:"when" json:"when"`
}
