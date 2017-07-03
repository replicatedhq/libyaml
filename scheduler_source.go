package libyaml

import "encoding/json"

type SchedulerContainerSource struct {
	*SourceContainerNative `yaml:"replicated,omitempty" json:"replicated,omitempty" validate:"omitempty,dive"`
	*SourceContainerSwarm  `yaml:"swarm,omitempty" json:"swarm,omitempty" validate:"omitempty,dive"`
	*SourceContainerK8s    `yaml:"kubernetes,omitempty" json:"kubernetes,omitempty" validate:"omitempty,dive"`
}

type SourceContainerNative struct {
	Component string `yaml:"component" json:"component" validate:"required,componentexists"`
	Container string `yaml:"container" json:"container" validate:"containerexists=Component"`
}

type SourceContainerSwarm struct {
	Service string `yaml:"service" json:"container" validate:"required"`
}

type SourceContainerK8s struct {
	Selectors map[string]string `yaml:"selectors" json:"selectors" validate:"required,dive,required"`
	Container string            `yaml:"container,omitempty" json:"container,omitempty"`
}

func (s *SchedulerContainerSource) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return s.unmarshal(unmarshal)
}

func (s *SchedulerContainerSource) UnmarshalJSON(data []byte) error {
	unmarshal := func(v interface{}) error {
		return json.Unmarshal(data, v)
	}
	return s.unmarshal(unmarshal)
}

// UnmarshalInline can be called inside of a parent's unmarshal function to allow
func UnmarshalInline(unmarshal func(interface{}) error, s *SchedulerContainerSource) error {
	var native SourceContainerNative
	if err := unmarshal(&native); err != nil {
		return err
	}
	if native.Component != "" {
		s.SourceContainerNative = &native
		return nil
	}
	var swarm SourceContainerSwarm
	if err := unmarshal(&swarm); err != nil {
		return err
	}
	if swarm.Service != "" {
		s.SourceContainerSwarm = &swarm
		return nil
	}
	var k8s SourceContainerK8s
	if err := unmarshal(&k8s); err != nil {
		return err
	}
	// container is kinda ambiguous, should determine if selectors is required
	if k8s.Selectors != nil || k8s.Container != "" {
		s.SourceContainerK8s = &k8s
		return nil
	}
	return nil
}

type schedulerContainerSourceInternal struct {
	*SourceContainerNative `yaml:"replicated" json:"replicated"`
	*SourceContainerSwarm  `yaml:"swarm" json:"swarm"`
	*SourceContainerK8s    `yaml:"kubernetes" json:"kubernetes"`
}

func (s *SchedulerContainerSource) unmarshal(unmarshal func(interface{}) error) error {
	var internal schedulerContainerSourceInternal
	if err := unmarshal(&internal); err != nil {
		return err
	}
	if internal.SourceContainerNative != nil || internal.SourceContainerSwarm != nil || internal.SourceContainerK8s != nil {
		s.SourceContainerNative = internal.SourceContainerNative
		s.SourceContainerSwarm = internal.SourceContainerSwarm
		s.SourceContainerK8s = internal.SourceContainerK8s
		return nil
	}
	return UnmarshalInline(unmarshal, s)
}
