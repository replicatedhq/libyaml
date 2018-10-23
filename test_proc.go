package libyaml

type TestProc struct {
	DisplayName   string                `yaml:"display_name" json:"display_name"`
	Command       string                `yaml:"command" json:"command"`
	CustomCommand CustomCommand         `yaml:"custom_command" json:"custom_command"`
	Results       []CustomCommandResult `yaml:"results" json:"results" validate:"dive"`
	Timeout       uint                  `yaml:"timeout" json:"timeout"`
	ArgFields     []string              `yaml:"arg_fields" json:"arg_fields"`
	Args          []string              `yaml:"args" json:"args"`
	RunOnSave     string                `yaml:"run_on_save" json:"run_on_save"`
}
