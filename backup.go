package libyaml

type Backup struct {
	Enabled         string `yaml:"enabled" json:"enabled"`
	PauseAll        bool   `yaml:"pause_all" json:"pause_all"` // deprecated
	PauseContainers string `yaml:"pause_containers" json:"pause_containers"`
	Script          string `yaml:"script" json:"script"`
}
