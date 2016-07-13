package libyaml_test

import (
	"testing"

	"github.com/replicatedhq/libyaml"

	"github.com/andreychernih/yaml"
)

var (
	k8sSpec = `apiVersion: v1
kind: Pod
metadata:
  name: hello-world
spec:  # specification of the pod's contents
  restartPolicy: Never
  containers:
  - name: hello
    image: "ubuntu:14.04"
    command: ["/bin/echo", "hello", "world"]`
)

func TestK8sUnmarshalYAML(t *testing.T) {
	var c libyaml.K8sConfig
	if err := yaml.Unmarshal([]byte(k8sSpec), &c); err != nil {
		t.Fatal(err)
	}

	var m map[string]interface{}
	if err := yaml.Unmarshal(c, &m); err != nil {
		t.Fatal(err)
	}

	testK8sYAML(t, m)
}

func TestK8sMarshalYAML(t *testing.T) {
	var c libyaml.K8sConfig = []byte(k8sSpec)
	b, err := yaml.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}

	var m map[string]interface{}
	if err := yaml.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}

	testK8sYAML(t, m)
}

func testK8sYAML(t *testing.T, m map[string]interface{}) {
	if val, ok := m["apiVersion"]; !ok {
		t.Errorf("%q not set", "apiVersion")
	} else if val, _ := val.(string); val != "v1" {
		t.Errorf("expecting %q == %q, got %q", "apiVersion", "v1", m["apiVersion"])
	}

	if val, ok := m["spec"].(map[interface{}]interface{})["containers"].([]interface{})[0].(map[interface{}]interface{})["image"]; !ok {
		t.Errorf("%q not set", "spec.containers[0].image")
	} else if val, _ := val.(string); val != "ubuntu:14.04" {
		t.Errorf("expecting %q == %q, got %q", "spec.containers[0].image", "ubuntu:14.04", val)
	}
}
