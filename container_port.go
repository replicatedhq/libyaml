package libyaml

type ContainerPort struct {
	PrivatePort string `yaml:"private_port" json:"private_port" validate:"required"`
	PublicPort  string `yaml:"public_port" json:"public_port" validate:"required"`
	Interface   string `yaml:"interface" json:"interface"`
	PortType    string `yaml:"port_type" json:"port_type"`
	When        string `yaml:"when" json:"when"`
}
