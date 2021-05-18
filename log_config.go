package libyaml

type LogOptions struct {
	MaxSize  string `yaml:"max_size,omitempty" json:"max_size,omitempty"`
	MaxFiles string `yaml:"max_files,omitempty" json:"max_files,omitempty"`
}

type LogConfig struct {
	Type   string            `yaml:"type" json:"type"`
	Config map[string]string `yaml:"config,omitempty" json:"config,omitempty"`
}
