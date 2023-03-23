package libyaml

type ConfigCommand struct {
	Name           string   `yaml:"name" json:"name"`
	Cmd            string   `yaml:"cmd" json:"cmd"`
	Args           []string `yaml:"args" json:"args"`
	ReadonlyRootfs bool     `yaml:"readonly_rootfs,omitempty" json:"readonly_rootfs,omitempty"`
}
