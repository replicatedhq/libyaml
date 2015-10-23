package libyaml

import (
	"testing"

	"github.com/andreychernih/yaml"
)

func TestComponentUnmarshalYAML(t *testing.T) {
	s := `name: test
tags: []
conflicts: []
cluster: false
cluster_host_count:
  min: 0
  threshold_healthy: 0
containers: []`

	var c Component
	if err := yaml.Unmarshal([]byte(s), &c); err != nil {
		t.Fatal(err)
	}

	if c.Name != "test" {
		t.Errorf("expecting \"Component.Name\" == \"test\", got \"%s\"", c.Name)
	}
	if c.Cluster {
		t.Error("expecting \"Component.Cluster\" to be false")
	}
	if c.ClusterHostCount.Min != 0 {
		t.Errorf("expecting \"Component.ClusterHostCount.Min\" == 0, got \"%d\"", c.ClusterHostCount.Min)
	}
	if c.ClusterHostCount.ThresholdHealthy != 0 {
		t.Errorf("expecting \"Component.ClusterHostCount.ThresholdHealthy\" == 0, got \"%d\"", c.ClusterHostCount.ThresholdHealthy)
	}
}

func TestComponentUnmarshalYAMLCluster(t *testing.T) {
	s := `name: test
tags: []
conflicts: []
cluster: true
containers: []`

	var c Component
	if err := yaml.Unmarshal([]byte(s), &c); err != nil {
		t.Fatal(err)
	}

	if c.Name != "test" {
		t.Errorf("expecting \"Component.Name\" == \"test\", got \"%s\"", c.Name)
	}
	if !c.Cluster {
		t.Error("expecting \"Component.Cluster\" to be true")
	}
	if c.ClusterHostCount.Min != 0 {
		t.Errorf("expecting \"Component.ClusterHostCount.Min\" == 0, got \"%d\"", c.ClusterHostCount.Min)
	}
	if c.ClusterHostCount.ThresholdHealthy != 0 {
		t.Errorf("expecting \"Component.ClusterHostCount.ThresholdHealthy\" == 0, got \"%d\"", c.ClusterHostCount.ThresholdHealthy)
	}
}

func TestComponentMarshalYAML(t *testing.T) {
	s := `name: test
tags: []
conflicts: []
cluster: false
cluster_host_count:
  min: 0
  threshold_healthy: 0
containers: []
`

	c := Component{
		Name:    "test",
		Cluster: false,
	}

	b, err := yaml.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != s {
		t.Errorf("unexpected marshalled YAML,\nexpecting\n%s\ngot\n%s", s, string(b))
	}
}

func TestComponentMarshalYAMLCluster(t *testing.T) {
	s := `name: test
tags: []
conflicts: []
cluster: true
cluster_host_count:
  min: 0
  threshold_healthy: 0
containers: []
`

	c := Component{
		Name:    "test",
		Cluster: true,
	}

	b, err := yaml.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != s {
		t.Errorf("unexpected marshalled YAML,\nexpecting\n%s\ngot\n%s", s, string(b))
	}
}
