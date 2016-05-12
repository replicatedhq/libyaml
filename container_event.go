package libyaml

type ContainerEvent struct {
	Name          string                        `yaml:"name" json:"name"`
	Trigger       string                        `yaml:"trigger" json:"trigger"`
	Data          string                        `yaml:"data" json:"data"`
	Args          []string                      `yaml:"args" json:"args"`
	Subscriptions []*ContainerEventSubscription `yaml:"subscriptions" json:"subscriptions" validate:"dive,exists"`
}
