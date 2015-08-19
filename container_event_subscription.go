package libyaml

type ContainerEventSubscription struct {
	ComponentName string `yaml:"component" json:"component"`
	ContainerName string `yaml:"container" json:"container"`
	Action        string `yaml:"action" json:"action"`
}
