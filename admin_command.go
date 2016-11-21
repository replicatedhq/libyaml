package libyaml

type AdminCommand struct {
	Alias     string       `yaml:"alias" json:"alias" validate:"shellalias"`
	Command   []string     `yaml:"command" json:"command"`
	Timeout   uint         `yaml:"timeout" json:"timeout"`
	RunType   string       `yaml:"run_type" json:"run_type"`
	Component string       `yaml:"component" json:"component" validate:"componentexists"`
	Image     CommandImage `yaml:"image" json:"image"` // TODO: validate exists
	When      string       `yaml:"when" json:"when"`
}

type CommandImage struct {
	Name    string `yaml:"image_name" json:"image_name"`
	Version string `yaml:"version" json:"version"`
}
