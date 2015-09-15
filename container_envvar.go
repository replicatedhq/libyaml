package libyaml

type ContainerEnvVar struct {
	Name                  string `yaml:"name" json:"name"`
	StaticVal             string `yaml:"static_val" json:"static_val"`
	IsExcludedFromSupport string `yaml:"is_excluded_from_support" json:"is_excluded_from_support"`
}
