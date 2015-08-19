package libyaml

type Backup struct {
	Enabled  string `yaml:"enabled" json:"enabled"`
	PauseAll bool   `yaml:"pause_all" json:"pause_all"`
	Script   string `yaml:"script" json:"script"`
}
