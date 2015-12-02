package libyaml

type ContainerSupportCommand struct {
	Filename string   `yaml:"filename" json:"filename" validate:"nonzero"`
	Command  []string `yaml:"command" json:"command"`
}
