package libyaml

type CustomRequirement struct {
	ID         string             `yaml:"id" json:"id" validate:"required,customrequirementidunique"`
	Message    Message            `yaml:"message" json:"message"`
	Details    *Message           `yaml:"details,omitempty" json:"details,omitempty"`
	When       BoolString         `yaml:"when,omitempty" json:"when,omitempty"`
	Results    []CustomResult     `yaml:"results" json:"results" validate:"required,min=1,dive"`
	Command    CustomCommand      `yaml:"command" json:"command"`
	Kubernetes *KubernetesCommand `yaml: "kubernetes,omitempty" json:"kubernetes,omitempty"`
}

type KubernetesCommand struct {
	Global      *bool                   `yaml:"global,omitempty" json:"global,omitempty"`
	Tolerations *[]KubernetesToleration `yaml:"tolerations,omitempty" json:"tolerations,omitempty"`
	PodName     string                  `yaml:"pod_name" json:"pod_name"`
}

type KubernetesToleration struct {
	Key      string `yaml:"key" json:"key"`
	Operator string `yaml:"operator" json:"operator"`
	Value    string `yaml:"value" json:"value"`
	Effect   string `yaml:"effect" json:"effect"`
}

type CustomResult struct {
	Status    string           `yaml:"status" json:"status" validate:"required"`
	Message   Message          `yaml:"message" json:"message"`
	Condition *CustomCondition `yaml:"condition,omitempty" json:"condition,omitempty"`
}

type CustomCondition struct {
	Error      bool       `yaml:"error,omitempty" json:"error,omitempty"`
	StatusCode *int       `yaml:"status_code,omitempty" json:"status_code,omitempty"`
	BoolExpr   BoolString `yaml:"bool_expr,omitempty" json:"bool_expr,omitempty"`
}

type CustomCommand struct {
	ID      string      `yaml:"id" json:"id" validate:"required"`
	Timeout int         `yaml:"timeout,omitempty" json:"timeout,omitempty"`
	Data    interface{} `yaml:"data,omitempty" json:"data,omitempty"`
}
