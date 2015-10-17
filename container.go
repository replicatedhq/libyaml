package libyaml

type Container struct {
	Source               string                        `yaml:"source" json:"source" validate:"required,externalregistryexists"`
	ImageName            string                        `yaml:"image_name" json:"image_name" validate:"required"`
	DisplayName          string                        `yaml:"display_name" json:"display_name"`
	Version              string                        `yaml:"version" json:"version" validate:"required"`
	Privileged           bool                          `yaml:"privileged" json:"privileged"`
	Hostname             string                        `yaml:"hostname" json:"hostname"`
	Cmd                  string                        `yaml:"cmd" json:"cmd"`
	Ephemeral            string                        `yaml:"ephemeral" json:"ephemeral"`
	SuppressRestart      []string                      `yaml:"suppress_restart" json:"suppress_restart"`
	Cluster              bool                          `yaml:"cluster" json:"cluster"`
	Restart              *ContainerRestartPolicy       `yaml:"restart" json:"restart"`
	ClusterInstanceCount ContainerClusterInstanceCount `yaml:"cluster_instance_count" json:"cluster_instance_count"`
	PublishEvents        []*ContainerEvent             `yaml:"publish_events" json:"publish_events" validate:"dive"`
	SubscribedEvents     []map[string]interface{}      `yaml:"-" json:"-"`
	ConfigFiles          []*ContainerConfigFile        `yaml:"config_files" json:"config_files" validate:"dive"`
	CustomerFiles        []*ContainerCustomerFile      `yaml:"customer_files" json:"customer_files" validate:"dive"`
	EnvVars              []*ContainerEnvVar            `yaml:"env_vars" json:"env_vars" validate:"dive"`
	Ports                []*ContainerPort              `yaml:"ports" json:"ports" validate:"dive"`
	Volumes              []*ContainerVolume            `yaml:"volumes" json:"volumes" validate:"dive"`
	ExtraHosts           []*ContainerExtraHost         `yaml:"extra_hosts" json:"hosts" validate:"dive"`
	SupportFiles         []*ContainerSupportFile       `yaml:"support_files" json:"support_files" validate:"dive"`
	SupportCommands      []*ContainerSupportCommand    `yaml:"support_commands" json:"support_commands" validate:"dive"`
	When                 string                        `yaml:"when" json:"when"`
}

type ContainerRestartPolicy struct {
	Policy string `yaml:"policy" json:"policy"`
	Max    uint   `yaml:"max" json:"max"`
}

type ContainerClusterInstanceCount struct {
	Initial           minInt1 `yaml:"initial" json:"initial"`
	Max               uint    `yaml:"max,omitempty" json:"max"` // 0 == unlimited
	ThresholdHealthy  uint    `yaml:"threshold_healthy" json:"threshold_healthy"`
	ThresholdDegraded uint    `yaml:"threshold_degraded,omitempty" json:"threshold_degraded"` // 0 == no degraded state
}

func (c *Container) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var m marshallerContainer
	if err := unmarshal(&m); err != nil {
		return err
	}
	m.decode(c)

	if c.Cluster {
		if c.ClusterInstanceCount.Initial == 0 {
			c.ClusterInstanceCount.Initial = 1
		}
	}

	return nil
}

func (c Container) MarshalYAML() (interface{}, error) {
	if !c.Cluster {
		m := nonclusterableContainer{}
		m.encode(c)
		return m, nil
	}

	m := marshallerContainer{}
	m.encode(c)
	return m, nil
}

type marshallerContainer Container

func (m *marshallerContainer) encode(c Container) {
	m.Source = c.Source
	m.ImageName = c.ImageName
	m.DisplayName = c.DisplayName
	m.Version = c.Version
	m.Privileged = c.Privileged
	m.Hostname = c.Hostname
	m.Cmd = c.Cmd
	m.Ephemeral = c.Ephemeral
	m.SuppressRestart = c.SuppressRestart
	m.Cluster = c.Cluster
	m.Restart = c.Restart
	m.ClusterInstanceCount = c.ClusterInstanceCount
	m.PublishEvents = c.PublishEvents
	m.SubscribedEvents = c.SubscribedEvents
	m.ConfigFiles = c.ConfigFiles
	m.CustomerFiles = c.CustomerFiles
	m.EnvVars = c.EnvVars
	m.Ports = c.Ports
	m.Volumes = c.Volumes
	m.ExtraHosts = c.ExtraHosts
	m.SupportFiles = c.SupportFiles
	m.SupportCommands = c.SupportCommands
	m.When = c.When
}

func (m marshallerContainer) decode(c *Container) {
	c.Source = m.Source
	c.ImageName = m.ImageName
	c.DisplayName = m.DisplayName
	c.Version = m.Version
	c.Privileged = m.Privileged
	c.Hostname = m.Hostname
	c.Cmd = m.Cmd
	c.Ephemeral = m.Ephemeral
	c.SuppressRestart = m.SuppressRestart
	c.Cluster = m.Cluster
	c.Restart = m.Restart
	c.ClusterInstanceCount = m.ClusterInstanceCount
	c.PublishEvents = m.PublishEvents
	c.SubscribedEvents = m.SubscribedEvents
	c.ConfigFiles = m.ConfigFiles
	c.CustomerFiles = m.CustomerFiles
	c.EnvVars = m.EnvVars
	c.Ports = m.Ports
	c.Volumes = m.Volumes
	c.ExtraHosts = m.ExtraHosts
	c.SupportFiles = m.SupportFiles
	c.SupportCommands = m.SupportCommands
	c.When = m.When
}

type nonclusterableContainer struct {
	Source           string                     `yaml:"source" json:"source" validate:"required,externalregistryexists"`
	ImageName        string                     `yaml:"image_name" json:"image_name" validate:"required"`
	DisplayName      string                     `yaml:"display_name" json:"display_name"`
	Version          string                     `yaml:"version" json:"version" validate:"required"`
	Privileged       bool                       `yaml:"privileged" json:"privileged"`
	Hostname         string                     `yaml:"hostname" json:"hostname"`
	Cmd              string                     `yaml:"cmd" json:"cmd"`
	Ephemeral        string                     `yaml:"ephemeral" json:"ephemeral"`
	SuppressRestart  []string                   `yaml:"suppress_restart" json:"suppress_restart"`
	Cluster          bool                       `yaml:"cluster" json:"cluster"`
	Restart          *ContainerRestartPolicy    `yaml:"restart" json:"restart"`
	PublishEvents    []*ContainerEvent          `yaml:"publish_events" json:"publish_events" validate:"dive"`
	SubscribedEvents []map[string]interface{}   `yaml:"-" json:"-"`
	ConfigFiles      []*ContainerConfigFile     `yaml:"config_files" json:"config_files" validate:"dive"`
	CustomerFiles    []*ContainerCustomerFile   `yaml:"customer_files" json:"customer_files" validate:"dive"`
	EnvVars          []*ContainerEnvVar         `yaml:"env_vars" json:"env_vars" validate:"dive"`
	Ports            []*ContainerPort           `yaml:"ports" json:"ports" validate:"dive"`
	Volumes          []*ContainerVolume         `yaml:"volumes" json:"volumes" validate:"dive"`
	ExtraHosts       []*ContainerExtraHost      `yaml:"extra_hosts" json:"hosts" validate:"dive"`
	SupportFiles     []*ContainerSupportFile    `yaml:"support_files" json:"support_files" validate:"dive"`
	SupportCommands  []*ContainerSupportCommand `yaml:"support_commands" json:"support_commands" validate:"dive"`
	When             string                     `yaml:"when" json:"when"`
}

func (m *nonclusterableContainer) encode(c Container) {
	m.Source = c.Source
	m.ImageName = c.ImageName
	m.DisplayName = c.DisplayName
	m.Version = c.Version
	m.Privileged = c.Privileged
	m.Hostname = c.Hostname
	m.Cmd = c.Cmd
	m.Ephemeral = c.Ephemeral
	m.SuppressRestart = c.SuppressRestart
	m.Cluster = false
	m.Restart = c.Restart
	m.PublishEvents = c.PublishEvents
	m.SubscribedEvents = c.SubscribedEvents
	m.ConfigFiles = c.ConfigFiles
	m.CustomerFiles = c.CustomerFiles
	m.EnvVars = c.EnvVars
	m.Ports = c.Ports
	m.Volumes = c.Volumes
	m.ExtraHosts = c.ExtraHosts
	m.SupportFiles = c.SupportFiles
	m.SupportCommands = c.SupportCommands
	m.When = c.When
}
