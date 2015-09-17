package libyaml

type Monitors struct {
	Cpuacct []string `yaml:"cpuacct" validate:"dive,componentcontainer,componentexists,containerexists"`
	Memory  []string `yaml:"memory" validate:"dive,componentcontainer,componentexists,containerexists"`
}
