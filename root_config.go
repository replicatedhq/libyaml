package libyaml

type RootConfig struct {
	ApiVersion             string           `yaml:"replicated_api_version" json:"replicated_api_version" validate:"required,apiversion"`
	Name                   string           `yaml:"name" json:"name"`
	Version                string           `yaml:"version" json:"version"`
	ReleaseNotes           string           `yaml:"release_notes" json:"release_notes"`
	ConsoleSupportMarkdown string           `yaml:"console_support_markdown" json:"console_support_markdown"`
	Properties             Properties       `yaml:"properties" json:"properties"`
	Identity               Identity         `yaml:"identity" json:"identity"`
	State                  State            `yaml:"state" json:"state"`
	Backup                 Backup           `yaml:"backup" json:"backup"`
	Monitors               Monitors         `yaml:"monitors" json:"monitors"`
	HostRequirements       HostRequirements `yaml:"host_requirements" json:"host_requirements"`
	ConfigCommands         []*ConfigCommand `yaml:"cmds" json:"cmds" validate:"dive,exists"`
	ConfigGroups           []*ConfigGroup   `yaml:"config" json:"config" validate:"dive,exists"`
	AdminCommands          []*AdminCommand  `yaml:"admin_commands" json:"admin_commands" validate:"dive,exists"`
	CustomMetrics          []*CustomMetric  `yaml:"custom_metrics" json:"custom_metrics" validate:"dive"`
	Graphite               Graphite         `yaml:"graphite" json:"graphite" validate:"dive"`

	Components []*Component `yaml:"components" json:"components" validate:"dive,exists"` // replicated scheduler config
	K8s        *K8s         `yaml:"kubernetes" json:"kubernetes"`
}

const DEFAULT_APP_CONFIG = `---
replicated_api_version: "1.3.2"
name: "%s"
properties:
  app_url: ""
  logo_url: http://www.replicated.com/images/logo.png
  console_title: Your App Name
backup:
  enabled: false
monitors:
  cpuacct: []
  memory: []
components: []
config: []`
