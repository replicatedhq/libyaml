package libyaml

type AdminCommand struct {
	Command     []string     `yaml:"command" json:"command"`
	Alias       string       `yaml:"alias" json:"alias"`
	Timeout     uint         `yaml:"timeout" json:"timeout"`
	RunType     string       `yaml:"run_type" json:"run_type"`
	Interactive bool         `yaml:"interactive" json:"interactive"`
	Component   string       `yaml:"component" json:"component"`
	Image       CommandImage `yaml:"image" json:"image"`
}

type CommandImage struct {
	Name    string `yaml:"image_name" json:"image_name"`
	Version string `yaml:"version" json:"version"`
}
