package libyaml_test

import (
	"reflect"
	"testing"

	. "github.com/replicatedhq/libyaml"

	yaml "gopkg.in/yaml.v2"
)

func TestUnmarshalSchedulerContainerSource(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		in := `
replicated:
  component: DB
  container: Redis
swarm:
  service: DB
kubernetes:
  selectors:
    app: redis
    role: master
  container: master`
		out := SchedulerContainerSource{}
		err := yaml.Unmarshal([]byte(in), &out)
		if err != nil {
			t.Fatal(err)
		}
		if "DB" != out.SourceContainerNative.Component {
			t.Errorf("expected:\n%s\nactual:\n%s", "DB", out.SourceContainerNative.Component)
		}
		if "Redis" != out.SourceContainerNative.Container {
			t.Errorf("expected:\n%s\nactual:\n%s", "Redis", out.SourceContainerNative.Container)
		}
		if "DB" != out.SourceContainerSwarm.Service {
			t.Errorf("expected:\n%s\nactual:\n%s", "DB", out.SourceContainerSwarm.Service)
		}
		expected := map[string]string{"app": "redis", "role": "master"}
		if !reflect.DeepEqual(expected, out.SourceContainerK8s.Selectors) {
			t.Errorf("expected:\n%s\nactual:\n%s", "DB", out.SourceContainerK8s.Selectors)
		}
		if "master" != out.SourceContainerK8s.Container {
			t.Errorf("expected:\n%s\nactual:\n%s", "master", out.SourceContainerK8s.Container)
		}
	})
}

func TestUnmarshalSchedulerContainerSourceNative(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		in := `
replicated:
  component: DB
  container: Redis`
		out := SchedulerContainerSource{}
		err := yaml.Unmarshal([]byte(in), &out)
		if err != nil {
			t.Fatal(err)
		}
		if "DB" != out.SourceContainerNative.Component {
			t.Errorf("expected:\n%s\nactual:\n%s", "DB", out.SourceContainerNative.Component)
		}
		if "Redis" != out.SourceContainerNative.Container {
			t.Errorf("expected:\n%s\nactual:\n%s", "Redis", out.SourceContainerNative.Container)
		}
		if nil != out.SourceContainerSwarm {
			t.Errorf("expected SourceContainerSwarm <nil>")
		}
		if nil != out.SourceContainerK8s {
			t.Errorf("expected SourceContainerK8s <nil>")
		}
	})

	t.Run("valid inline", func(t *testing.T) {
		in := `
component: DB
container: Redis`
		out := SchedulerContainerSource{}
		err := yaml.Unmarshal([]byte(in), &out)
		if err != nil {
			t.Fatal(err)
		}
		if "DB" != out.SourceContainerNative.Component {
			t.Errorf("expected:\n%s\nactual:\n%s", "DB", out.SourceContainerNative.Component)
		}
		if "Redis" != out.SourceContainerNative.Container {
			t.Errorf("expected:\n%s\nactual:\n%s", "Redis", out.SourceContainerNative.Container)
		}
		if nil != out.SourceContainerSwarm {
			t.Errorf("expected SourceContainerSwarm <nil>")
		}
		if nil != out.SourceContainerK8s {
			t.Errorf("expected SourceContainerK8s <nil>")
		}
	})
}

func TestUnmarshalSchedulerContainerSourceSwarm(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		in := `
swarm:
  service: DB`
		out := SchedulerContainerSource{}
		err := yaml.Unmarshal([]byte(in), &out)
		if err != nil {
			t.Fatal(err)
		}
		if "DB" != out.SourceContainerSwarm.Service {
			t.Errorf("expected:\n%s\nactual:\n%s", "DB", out.SourceContainerSwarm.Service)
		}
		if nil != out.SourceContainerNative {
			t.Errorf("expected SourceContainerNative <nil>")
		}
		if nil != out.SourceContainerK8s {
			t.Errorf("expected SourceContainerK8s <nil>")
		}
	})

	t.Run("valid inline", func(t *testing.T) {
		in := `
service: DB`
		out := SchedulerContainerSource{}
		err := yaml.Unmarshal([]byte(in), &out)
		if err != nil {
			t.Fatal(err)
		}
		if "DB" != out.SourceContainerSwarm.Service {
			t.Errorf("expected:\n%s\nactual:\n%s", "DB", out.SourceContainerSwarm.Service)
		}
		if nil != out.SourceContainerNative {
			t.Errorf("expected SourceContainerNative <nil>")
		}
		if nil != out.SourceContainerK8s {
			t.Errorf("expected SourceContainerK8s <nil>")
		}
	})
}

func TestUnmarshalSchedulerContainerSourceK8s(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		in := `
kubernetes:
  selectors:
    app: redis
    role: master
  container: master`
		out := SchedulerContainerSource{}
		err := yaml.Unmarshal([]byte(in), &out)
		if err != nil {
			t.Fatal(err)
		}
		expected := map[string]string{"app": "redis", "role": "master"}
		if !reflect.DeepEqual(expected, out.SourceContainerK8s.Selectors) {
			t.Errorf("expected:\n%s\nactual:\n%s", "DB", out.SourceContainerK8s.Selectors)
		}
		if "master" != out.SourceContainerK8s.Container {
			t.Errorf("expected:\n%s\nactual:\n%s", "master", out.SourceContainerK8s.Container)
		}
		if nil != out.SourceContainerNative {
			t.Errorf("expected SourceContainerNative <nil>")
		}
		if nil != out.SourceContainerSwarm {
			t.Errorf("expected SourceContainerK8s <nil>")
		}
	})

	t.Run("valid inline", func(t *testing.T) {
		in := `
selectors:
  app: redis
  role: master
container: master`
		out := SchedulerContainerSource{}
		err := yaml.Unmarshal([]byte(in), &out)
		if err != nil {
			t.Fatal(err)
		}
		expected := map[string]string{"app": "redis", "role": "master"}
		if !reflect.DeepEqual(expected, out.SourceContainerK8s.Selectors) {
			t.Errorf("expected:\n%s\nactual:\n%s", "DB", out.SourceContainerK8s.Selectors)
		}
		if "master" != out.SourceContainerK8s.Container {
			t.Errorf("expected:\n%s\nactual:\n%s", "master", out.SourceContainerK8s.Container)
		}
		if nil != out.SourceContainerNative {
			t.Errorf("expected SourceContainerNative <nil>")
		}
		if nil != out.SourceContainerSwarm {
			t.Errorf("expected SourceContainerK8s <nil>")
		}
	})
}
