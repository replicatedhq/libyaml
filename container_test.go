package libyaml

import (
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func TestContainerUnmarshalYAML(t *testing.T) {
	s := `source: public
image_name: test
display_name: Test Container
version: ""
privileged: false
hostname: ""
cmd: ""
cluster: false
publish_events: []
config_files: []
customer_files: []
env_vars: []
ports: []
volumes: []
support_files: []
support_commands: []`

	var c Container
	if err := yaml.Unmarshal([]byte(s), &c); err != nil {
		t.Fatal(err)
	}

	if c.Source != "public" {
		t.Errorf("expecting \"Container.Source\" == \"public\", got \"%s\"", c.Source)
	}
	if c.ImageName != "test" {
		t.Errorf("expecting \"Container.ImageName\" == \"test\", got \"%s\"", c.ImageName)
	}
	if c.DisplayName != "Test Container" {
		t.Errorf("expecting \"Container.DisplayName\" == \"Test Container\", got \"%s\"", c.ImageName)
	}
	if c.Cluster {
		t.Error("expecting \"Container.Cluster\" to be false")
	}
	if c.ClusterInstanceCount.Initial != 0 {
		t.Errorf("expecting \"Container.ClusterInstanceCount.Initial\" == 0, got \"%d\"", c.ClusterInstanceCount.Initial)
	}
	if c.ClusterInstanceCount.ThresholdHealthy != 0 {
		t.Errorf("expecting \"Container.ClusterInstanceCount.ThresholdHealthy\" == 0, got \"%d\"", c.ClusterInstanceCount.ThresholdHealthy)
	}
}

func TestContainerUnmarshalYAMLCluster(t *testing.T) {
	s := `source: public
image_name: test
display_name: Test Container
version: ""
privileged: false
hostname: ""
cmd: ""
cluster: true
publish_events: []
config_files: []
customer_files: []
env_vars: []
ports: []
volumes: []
support_files: []
support_commands: []`

	var c Container
	if err := yaml.Unmarshal([]byte(s), &c); err != nil {
		t.Fatal(err)
	}

	if c.Source != "public" {
		t.Errorf("expecting \"Container.Source\" == \"public\", got \"%s\"", c.Source)
	}
	if c.ImageName != "test" {
		t.Errorf("expecting \"Container.ImageName\" == \"test\", got \"%s\"", c.ImageName)
	}
	if c.DisplayName != "Test Container" {
		t.Errorf("expecting \"Container.DisplayName\" == \"Test Container\", got \"%s\"", c.ImageName)
	}
	if !c.Cluster {
		t.Error("expecting \"Container.Cluster\" to be true")
	}
	if c.ClusterInstanceCount.Initial != 1 {
		t.Errorf("expecting \"Container.ClusterInstanceCount.Initial\" == 1, got \"%d\"", c.ClusterInstanceCount.Initial)
	}
	if c.ClusterInstanceCount.ThresholdHealthy != 0 {
		t.Errorf("expecting \"Container.ClusterInstanceCount.ThresholdHealthy\" == 0, got \"%d\"", c.ClusterInstanceCount.ThresholdHealthy)
	}
}

func TestContainerMarshalYAML(t *testing.T) {
	s := `source: public
image_name: test
display_name: ""
version: ""
privileged: false
hostname: ""
cmd: ""
ephemeral: false
cluster: false
publish_events: []
config_files: []
customer_files: []
env_vars: []
ports: []
volumes: []
support_files: []
support_commands: []
`

	c := Container{
		Source:    "public",
		ImageName: "test",
		Cluster:   false,
	}

	b, err := yaml.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != s {
		t.Errorf("unexpected marshalled YAML,\nexpecting\n%s\ngot\n%s", s, string(b))
	}
}

func TestContainerMarshalYAMLCluster(t *testing.T) {
	s := `source: public
image_name: test
display_name: ""
version: ""
privileged: false
hostname: ""
cmd: ""
ephemeral: false
cluster: true
cluster_instance_count:
  initial: 1
  threshold_healthy: 0
publish_events: []
config_files: []
customer_files: []
env_vars: []
ports: []
volumes: []
support_files: []
support_commands: []
`

	c := Container{
		Source:    "public",
		ImageName: "test",
		Cluster:   true,
	}

	b, err := yaml.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != s {
		t.Errorf("unexpected marshalled YAML,\nexpecting\n%s\ngot\n%s", s, string(b))
	}
}
