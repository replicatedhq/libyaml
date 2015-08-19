package libyaml

type ContainerPort struct {
	PrivatePort string `yaml:"private_port" json:"private_port"`
	PublicPort  string `yaml:"public_port" json:"public_port"`
	Interface   string `yaml:"interface" json:"interface"`
	PortType    string `yaml:"port_type" json:"port_type"`
	When        string `yaml:"when" json:"when"`
}
