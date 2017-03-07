package libyaml

import "encoding/json"

var (
	AdminCommandRunTypeExec AdminCommandRunType = "exec"
)

type AdminCommand struct {
	// AdminCommandV2 api version >= 2.6.0
	AdminCommandV2 `yaml:",inline"`
	// AdminCommandV1 api version < 2.6.0
	AdminCommandV1 `yaml:",inline"`
}

type AdminCommandV2 struct {
	Alias   string              `yaml:"alias" json:"alias" validate:"required,shellalias"`
	Command []string            `yaml:"command,flow" json:"command" validate:"required"`
	Timeout uint                `yaml:"timeout,omitempty" json:"timeout,omitempty"`
	RunType AdminCommandRunType `yaml:"run_type,omitempty" json:"run_type,omitempty"` // default "exec"
	When    string              `yaml:"when,omitempty" json:"when,omitempty"`
	Source  AdminCommandSource  `yaml:"source" json:"source" validate:"required"`
}

type AdminCommandSource struct {
	Replicated *AdminCommandSourceReplicated `yaml:"replicated,omitempty" json:"replicated,omitempty" validate:"omitempty,dive"`
	Kubernetes *AdminCommandSourceKubernetes `yaml:"kubernetes,omitempty" json:"kubernetes,omitempty" validate:"omitempty,dive"`
}

type AdminCommandSourceReplicated struct {
	Component string `yaml:"component" json:"component" validate:"required,componentexists"`
	Container string `yaml:"container" json:"container" validate:"containerexists=Component"`
}

type AdminCommandSourceKubernetes struct {
	Selectors map[string]string `yaml:"selectors" json:"selectors" validate:"required,dive,required"`
	Container string            `yaml:"container,omitempty" json:"container,omitempty"`
}

type AdminCommandRunType string

type AdminCommandV1 struct { // deprecated
	Component string        `yaml:"component,omitempty" json:"component,omitempty" validate:"omitempty,componentexists"`
	Image     *CommandImage `yaml:"image,omitempty" json:"image,omitempty" validate:"omitempty,dive"`
}

type CommandImage struct {
	Name    string `yaml:"image_name" json:"image_name" validate:"required"`
	Version string `yaml:"version" json:"version"`
}

func (c *AdminCommand) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return c.unmarshal(unmarshal)
}

func (c *AdminCommand) UnmarshalJSON(data []byte) error {
	unmarshal := func(v interface{}) error {
		return json.Unmarshal(data, v)
	}
	return c.unmarshal(unmarshal)
}

func (c *AdminCommand) unmarshal(unmarshal func(interface{}) error) error {
	v2 := AdminCommandV2{}
	if err := unmarshal(&v2); err != nil {
		return err
	}
	c.AdminCommandV2 = v2

	v1 := AdminCommandV1{}
	if err := unmarshal(&v1); err != nil {
		return err
	}
	c.AdminCommandV1 = v1

	if c.Source.Replicated == nil {
		out := &AdminCommandSourceReplicated{}
		if err := unmarshal(out); err == nil && out.Component != "" {
			c.Source.Replicated = out
		}
	}
	if c.Source.Kubernetes == nil {
		out := &AdminCommandSourceKubernetes{}
		if err := unmarshal(out); err == nil && out.Selectors != nil {
			c.Source.Kubernetes = out
		}
	}

	// backwards compatibility
	if c.Source.Replicated != nil {
		if c.Image == nil {
			c.Image = &CommandImage{}
		}

		if c.Component == "" {
			c.Component = c.Source.Replicated.Component
		}

		if c.Source.Replicated.Container == "" {
			c.Source.Replicated.Container = c.Image.Name
		} else {
			c.Image.Name = c.Source.Replicated.Container
		}
	}

	return nil
}
