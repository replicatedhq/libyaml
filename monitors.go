package libyaml

type Monitors struct {
	Cpuacct []string `yaml:"cpuacct"`
	Memory  []string `yaml:"memory"`
}
