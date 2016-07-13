package libyaml_test

import (
	"testing"

	"github.com/replicatedhq/libyaml"

	"github.com/andreychernih/yaml"
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
logs:
  max_size: "100k"
  max_files: "5"
volumes: []
support_files: []
support_commands: []`

	var c libyaml.Container
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
	if c.LogOptions.MaxFiles != "5" {
		t.Errorf("expecting \"Container.MaxFiles\" == \"5\", got \"%s\"", c.LogOptions.MaxFiles)
	}
	if c.LogOptions.MaxSize != "100k" {
		t.Errorf("expecting \"Container.MaxSize\" == \"100k\", got \"%s\"", c.LogOptions.MaxSize)
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
logs:
  max_size: ""
  max_files: ""
volumes: []
support_files: []
support_commands: []`

	var c libyaml.Container
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
network_mode: ""
cpu_shares: ""
memory_limit: ""
memory_swap_limit: ""
allocate_tty: ""
security_cap_add: []
security_options: []
hostname: ""
cmd: ""
ephemeral: false
suppress_restart: []
cluster: false
restart: null
publish_events: []
config_files: []
customer_files: []
env_vars: []
ports: []
logs:
  max_size: 100k
  max_files: "5"
volumes: []
extra_hosts: []
support_files: []
support_commands: []
when: ""
`

	logReqs := libyaml.LogOptions{
		MaxSize:  "100k",
		MaxFiles: "5",
	}
	c := libyaml.Container{
		Source:     "public",
		ImageName:  "test",
		Cluster:    false,
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

func TestContainerMarshalYAMLCluster(t *testing.T) {
	s := `source: public
image_name: test
display_name: ""
version: ""
privileged: false
network_mode: ""
cpu_shares: ""
memory_limit: ""
memory_swap_limit: ""
allocate_tty: ""
security_cap_add: []
security_options: []
hostname: ""
cmd: ""
ephemeral: false
suppress_restart: []
cluster: true
restart: null
cluster_instance_count:
  initial: 1
  threshold_healthy: 0
publish_events: []
config_files: []
customer_files: []
env_vars: []
ports: []
logs:
  max_size: ""
  max_files: ""
volumes: []
extra_hosts: []
support_files: []
support_commands: []
when: ""
`

	c := libyaml.Container{
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
