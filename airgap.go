package libyaml

type Image struct {
	Source       string        `yaml:"source,omitempty" json:"source,omitempty" validate:"externalregistryexists"` // default public
	Name         string        `yaml:"name" json:"name" validate:"required"`
	Tag          string        `yaml:"tag,omitempty" json:"tag,omitempty"` // default latest
	ContentTrust *ContentTrust `yaml:"content_trust,omitempty" json:"content_trust,omitempty"`
}
