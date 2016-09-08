package libyaml_test

import (
	"testing"

	"github.com/replicatedhq/libyaml"

	"gopkg.in/yaml.v2"
)

func TestComponentUnmarshalYAML(t *testing.T) {
	s := `name: test
tags: []
conflicts: []
cluster: false
cluster_host_count:
  min: 0
  threshold_healthy: 0
logs:
  max_size: 100k
  max_files: 5
containers: []`

	var c libyaml.Component
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
	if c.LogOptions.MaxFiles != "5" {
		t.Errorf("expecting \"Component.MaxFiles\" == \"5\", got \"%s\"", c.LogOptions.MaxFiles)
	}
	if c.LogOptions.MaxSize != "100k" {
		t.Errorf("expecting \"Component.MaxSize\" == \"100k\", got \"%s\"", c.LogOptions.MaxSize)
	}
}

func TestComponentUnmarshalYAMLCluster(t *testing.T) {
	s := `name: test
tags: []
conflicts: []
cluster: true
containers: []`

	var c libyaml.Component
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
host_requirements:
  cpu_cores: 0
  cpu_mhz: 0
  memory: ""
  disk_space: ""
logs:
  max_size: ""
  max_files: ""
host_volumes: []
containers: []
`

	c := libyaml.Component{
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
host_requirements:
  cpu_cores: 0
  cpu_mhz: 0
  memory: ""
  disk_space: ""
logs:
  max_size: 100k
  max_files: "5"
host_volumes: []
containers: []
`

	logReqs := libyaml.LogOptions{
		MaxSize:  "100k",
		MaxFiles: "5",
	}
	c := libyaml.Component{
		Name:       "test",
		Cluster:    true,
		LogOptions: logReqs,
	}

	b, err := yaml.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != s {
		t.Errorf("unexpected marshalled YAML,\nexpecting\n%s\ngot\n%s", s, string(b))
	}
}
