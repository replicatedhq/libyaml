package libyaml

import "github.com/andreychernih/yaml"

type K8s struct {
	Config K8sConfig `yaml:"config" validate:"required"`
}

type K8sConfig []byte

func (c *K8sConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var m interface{}
	if err := unmarshal(&m); err != nil {
		return err
	}
	b, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	*c = K8sConfig(b)
	return nil
}

func (c K8sConfig) MarshalYAML() (interface{}, error) {
	var m interface{}
	err := yaml.Unmarshal(c, &m)
	return m, err
}
