package libyaml

type ContainerSupportCommand struct {
	Filename string   `yaml:"filename" json:"filename"`
	Command  []string `yaml:"command" json:"command"`
}
