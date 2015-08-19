package libyaml

type RootConfig struct {
	ApiVersion     string           `yaml:"replicated_api_version" json:"replicated_api_version"`
	Name           string           `yaml:"name" json:"name"`
	Version        string           `yaml:"version" json:"version"`
	ReleaseNotes   string           `yaml:"release_notes" json:"release_notes"`
	Properties     Properties       `yaml:"properties" json:"properties"`
	Identity       Identity         `yaml:"identity" json:"identity"`
	State          State            `yaml:"state" json:"state"`
	Backup         Backup           `yaml:"backup" json:"backup"`
	Monitors       Monitors         `yaml:"monitors" json:"monitors"`
	Components     []*Component     `yaml:"components" json:"components"`
	ConfigCommands []*ConfigCommand `yaml:"cmds" json:"cmds"`
	ConfigGroups   []*ConfigGroup   `yaml:"config" json:"config"`
	AdminCommands  []*AdminCommand  `yaml:"admin_commands" json:"admin_commands"`
}

const DEFAULT_APP_CONFIG = `
---
replicated_api_version: "1.0.0"
name: "%s"
version: ""
release_notes: ""
components: []
backup:
  enabled: false
  init: []
  run: []
cmds: []
config: []
`
