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
	// CustomRequirements api version >= 2.4.0
	CustomRequirements []CustomRequirement `yaml:"custom_requirements,omitempty" json:"custom_requirements,omitempty" validate:"dive"`
	ConfigCommands     []*ConfigCommand    `yaml:"cmds" json:"cmds" validate:"dive,exists"`
	ConfigGroups       []*ConfigGroup      `yaml:"config" json:"config" validate:"dive,exists"`
	AdminCommands      []*AdminCommand     `yaml:"admin_commands" json:"admin_commands" validate:"dive,exists"`
	CustomMetrics      []*CustomMetric     `yaml:"custom_metrics" json:"custom_metrics" validate:"dive"`
	Graphite           Graphite            `yaml:"graphite" json:"graphite" validate:"dive"`
	StatsD             StatsD              `yaml:"statsd" json:"statsd" validate:"dive"`
	Localization       *Localization       `yaml:"localization,omitempty" json:"localization,omitempty" validate:"omitempty,dive"`

	Components []*Component `yaml:"components" json:"components" validate:"dive,exists"` // replicated scheduler config
	K8s        *K8s         `yaml:"kubernetes" json:"kubernetes"`                        // this is deprecated, prefer multi-doc
}

const DEFAULT_APP_CONFIG = `---
replicated_api_version: 2.4.2
name: "%s"

#
# https://www.replicated.com/docs/packaging-an-application/application-properties
#
properties:
  app_url: http://{{repl ConfigOption "hostname" }}
  console_title: "%s"

#
# Settings screen
# https://www.replicated.com/docs/packaging-an-application/config-screen
#
config:
- name: hostname
  title: Hostname
  description: Ensure this domain name is routable on your network.
  items:
  - name: hostname
    title: Hostname
    value: '{{repl ConsoleSetting "tls.hostname"}}'
    type: text
    test_proc:
      display_name: Check DNS
      command: resolve_host

#
# Define how the application containers will be created and started
# https://www.replicated.com/docs/packaging-an-application/components-and-containers
#
components: []

#
# Documentation for additional features
# https://www.replicated.com/docs/packaging-an-application
#`
