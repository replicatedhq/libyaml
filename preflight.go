package libyaml

type CustomRequirement struct {
	ID      string         `yaml:"id" json:"id" validate:"required,customrequirementidunique"`
	Message Message        `yaml:"message" json:"message"`
	Details *Message       `yaml:"details,omitempty" json:"details,omitempty"`
	When    BoolString     `yaml:"when,omitempty" json:"when,omitempty"`
	Tags    []string       `yaml:"tags,omitempty" json:"tags,omitempty" validate:"dive,required"`
	Results []CustomResult `yaml:"results" json:"results" validate:"required,min=1,dive"`
	Command CustomCommand  `yaml:"command" json:"command"`
}

type CustomResult struct {
	Status    string           `yaml:"status" json:"status" validate:"required"`
	Message   interface{}      `yaml:"message" json:"message" validate:"required"`
	Condition *CustomCondition `yaml:"condition,omitempty" json:"condition,omitempty"`
}

type CustomCondition struct {
	Error      bool       `yaml:"error,omitempty" json:"error,omitempty"`
	StatusCode string     `yaml:"status_code,omitempty" json:"status_code,omitempty" validate:"omitempty,int"`
	BoolExpr   BoolString `yaml:"bool_expr,omitempty" json:"bool_expr,omitempty"`
}

type CustomCommand struct {
	ID           string        `yaml:"id" json:"id" validate:"required"`
	Source       string        `yaml:"source,omitempty" json:"source,omitempty" validate:"omitempty,externalregistryexists"`
	ImageName    string        `yaml:"image_name,omitempty" json:"image_name,omitempty"`
	Tag          string        `yaml:"tag,omitempty" json:"tag,omitempty"`
	Version      string        `yaml:"version,omitempty" json:"version,omitempty"` // alias of tag
	ContentTrust *ContentTrust `yaml:"content_trust,omitempty" json:"content_trust,omitempty"`
	Timeout      int           `yaml:"timeout,omitempty" json:"timeout,omitempty"`
	Data         interface{}   `yaml:"data,omitempty" json:"data,omitempty"`
}
