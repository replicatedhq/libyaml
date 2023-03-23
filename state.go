package libyaml

type State struct {
	Ready *StateEvent `yaml:"ready" json:"ready"`
}

type StateEvent struct {
	Command        string     `yaml:"command" json:"command"`
	Args           []string   `yaml:"args" json:"args"`
	Timeout        int        `yaml:"timeout" json:"timeout"`
	ReadonlyRootfs BoolString `yaml:"readonly_rootfs,omitempty" json:"readonly_rootfs,omitempty"`
}
