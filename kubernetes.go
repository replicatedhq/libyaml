package libyaml

type K8s struct {
	Config string `yaml:"config" validate:"required"`
}
