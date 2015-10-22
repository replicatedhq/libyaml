package libyaml

type ConfigItem struct {
	Name        string         `yaml:"name" json:"name"`
	Title       string         `yaml:"title" json:"title"`
	HelpText    string         `yaml:"help_text" json:"help_text"`
	Recommended bool           `yaml:"recommended" json:"recommended"`
	Default     string         `yaml:"default" json:"default"`
	Value       string         `yaml:"value" json:"value"`
	MultiValue  []string       `yaml:"multi_value" json:"multi_value"`
	DefaultCmd  *ConfigItemCmd `yaml:"default_cmd" json:"default_cmd"`
	ValueCmd    *ConfigItemCmd `yaml:"value_cmd" json:"value_cmd"`
	DataCmd     *ConfigItemCmd `yaml:"data_cmd" json:"data_cmd"`
	ReadOnly    bool           `yaml:"readonly" json:"readonly"`
	When        string         `yaml:"when" json:"when"`
	Type        string         `yaml:"type" json:"type"`
	Multiple    bool           `yaml:"multiple" json:"multiple"`
	Hidden      bool           `yaml:"hidden" json:"hidden"`
	Position    int            `yaml:"-" json:"-"`
	Affix       string         `yaml:"affix" json:"affix"`
	Required    bool           `yaml:"required" json:"required"`
	TestProc    *TestProc      `yaml:"test_proc" json:"test_proc"`
	Items       []*ConfigItem  `yaml:"items" json:"items"`
}

type ConfigItemCmd struct {
	Name    string `yaml:"name" json:"name"`
	ValueAt int    `yaml:"value_at" json:"value_at"`
}

type TestProc struct {
	DisplayName string   `yaml:"display_name" json:"display_name"`
	Command     string   `yaml:"command" json:"command"`
	Timeout     uint     `yaml:"timeout" json:"timeout"`
	ArgFields   []string `yaml:"arg_fields" json:"arg_fields"`
}
