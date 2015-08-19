package libyaml

type ContainerConfigFile struct {
	Filename string `yaml:"filename" json:"filename"`
	Contents string `yaml:"contents" json:"contents"`
	Source   string `yaml:"source" json:"source"`
	Owner    string `yaml:"owner" json:"owner"`
	Repo     string `yaml:"repo" json:"repo"`
	Path     string `yaml:"path" json:"path"`
	Ref      string `yaml:"ref" json:"ref"`
}
