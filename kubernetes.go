package libyaml

type K8s struct {
	Config       string          `yaml:"config"` // this is deprecated, prefer multi-doc
	Requirements K8sRequirements `yaml:"requirements,omitempty" json:"requirements,omitempty"`
	PVClaims     []K8sPVClaim    `yaml:"persistent_volume_claims,omitempty" json:"persistent_volume_claims,omitempty" validate:"dive"`
	SharedFS     K8sSharedFS     `yaml:"shared_fs,omitempty" json:"shared_fs,omitempty"`
}

type K8sRequirements struct {
	ServerVersion string     `yaml:"server_version,omitempty" json:"server_version,omitempty" validate:"omitempty,semverrange"`
	APIVersions   []string   `yaml:"api_versions,omitempty" json:"api_versions,omitempty" validate:"dive,required"`
	ClusterSize   UintString `yaml:"cluster_size,omitempty" json:"cluster_size,omitempty" validate:"omitempty,uint"`
	TotalCores    UintString `yaml:"total_cores,omitempty" json:"total_cores,omitempty" validate:"omitempty,uint"`
	TotalMemory   string     `yaml:"total_memory,omitempty" json:"total_memory,omitempty" validate:"omitempty,bytes|ram|quantity"`
}

type K8sPVClaim struct {
	Name        string   `yaml:"name" json:"name" validate:"required"`
	Storage     string   `yaml:"storage,omitempty" json:"storage,omitempty" validate:"omitempty,bytes|quantity"`
	AccessModes []string `yaml:"access_modes,omitempty" json:"access_modes,omitempty"`
}

type K8sSharedFS struct {
	Enabled    string   `yaml:"enabled,omitempty" json:"enabled,omitempty"`
	MountPaths []string `yaml:"mount_paths,omitempty" json:"mount_paths,omitempty"`
}
