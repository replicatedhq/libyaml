package libyaml

type RootK8s struct {
	VolumeClaims []K8sVolumeClaim `yaml:"volume_claims,omitempty" json:"volume_claims,omitempty"`
}

type K8sVolumeClaim struct {
	Name        string   `yaml:"name" json:"name" validate:"required"`
	Storage     string   `yaml:"storage,omitempty" json:"storage,omitempty" validate:"bytes"`
	AccessModes []string `yaml:"access_modes,omitempty" json:"access_modes,omitempty"`
}
