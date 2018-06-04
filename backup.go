package libyaml

type BackupStrategy struct {
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`

	// If Backup.Strategies is not an empty list, fields with the
	// same names in the Backup struct will be ignored.
	ExcludeAppData        string       `yaml:"exclude_app_data" json:"exclude_app_data"`
	ExcludeReplicatedData string       `yaml:"exclude_replicated_data" json:"exclude_replicated_data"`
	Enabled               string       `yaml:"enabled" json:"enabled"`
	PauseContainers       string       `yaml:"pause_containers" json:"pause_containers"`
	Script                string       `yaml:"script" json:"script"`
	Manual                string       `yaml:"manual" json:"manual"`
	Kubernetes            K8sBackups   `yaml:"kubernetes" json:"kubernetes"`
	Swarm                 SwarmBackups `yaml:"swarm" json:"swarm"`
}

type Backup struct {
	Enabled         string           `yaml:"enabled" json:"enabled"`
	Hidden          string           `yaml:"hidden" json:"hidden"`
	PauseAll        bool             `yaml:"pause_all" json:"pause_all"` // deprecated
	PauseContainers string           `yaml:"pause_containers" json:"pause_containers"`
	Script          string           `yaml:"script" json:"script"`
	Kubernetes      K8sBackups       `yaml:"kubernetes" json:"kubernetes"`
	Swarm           SwarmBackups     `yaml:"swarm" json:"swarm"`
	Strategies      []BackupStrategy `yaml:"strategies" json:"strategies"`
}

type K8sBackups struct {
	PVCNames []string `yaml:"pvc_names" json:"pvc_names"`
}

type SwarmBackups struct {
	Volumes []string `yaml:"volumes" json:"volumes"`
}
