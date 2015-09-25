package libyaml

type ContainerCustomerFile struct {
	Id       string `yaml:"name" json:"name"`
	Filename string `yaml:"filename" json:"filename"`
	When     string `yaml:"when" json:"when"`
}
