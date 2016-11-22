package libyaml_test

import (
	"strings"
	"testing"

	libyaml "github.com/replicatedhq/libyaml"
	validator "gopkg.in/go-playground/validator.v8"
	"gopkg.in/yaml.v2"
)

func TestTemplateFunctionPassesValidation(t *testing.T) {
	s := `
replicated_api_version: "1.0.0"
name: validate_templ_funcs
components:
- name: Elasticsearch
  containers:
  - source: public
    image_name: elasticsearch
    version: 1.5.0-3
    ports:
    - public_port: 9200
      private_port: 9200
  - source: public
    image_name: plugin-head
    version: 0.2.0
    env_vars:
    # Generic functions, always available
    - name: Now
      static_val: '{{repl Now}}'
    - name: Now
      static_val: '{{repl NowFmt "param"}}'
    - name: ToLower
      static_val: '{{repl ToLower "param"}}'
    - name: ToUpper
      static_val: '{{repl ToUpper "param"}}'
    - name: UrlEncode
      static_val: '{{repl UrlEncode "param"}}'
    - name: Base64Encode
      static_val: '{{repl Base64Encode "param"}}'
    - name: Base64Decode
      static_val: '{{repl Base64Decode "param"}}'
    - name: Split
      static_val: '{{repl Split "foo,bar,baz" ","}}'
    - name: Add
      static_val: '{{repl Add "100" "-1"}}'
    - name: Sub
      static_val: '{{repl Sub "1" "1"}}'
    - name: Mult
      static_val: '{{repl Mult "1" "1.5"}}'
    - name: Div
      static_val: '{{repl Div "10" "5"}}'

    # Always available
    - name: Pipeline
      static_val: '{{repl if ConfigOptionEquals "var" "val"}} T1 {{repl else}} T0 {{repl end}}'
    - name: ConfigOption
      static_val: '{{repl ConfigOption "fieldName"}}'
    - name: ConfigOptionIndex
      static_val: '{{repl ConfigOptionIndex "fieldName" 34}}'
    - name: ConfigOptionData
      static_val: '{{repl ConfigOptionData "filename"}}'
    - name: ConfigOptionEquals
      static_val: '{{repl ConfigOptionEquals "a" "b"}}'
    - name: ConfigOptionNotEquals
      static_val: '{{repl ConfigOptionNotEquals "a" "b"}}'
    - name: LicenseFieldValue
      static_val: '{{repl LicenseFieldValue "fieldName"}}'
    - name: LicenseProperty
      static_val: '{{repl LicenseProperty "propertyName"}}'
    - name: LdapCopyAuthFrom
      static_val: '{{repl LdapCopyAuthFrom "Hostname"}}'
    - name: ConsoleSetting
      static_val: '{{repl ConsoleSetting "tls.key.name"}}'
    - name: ConsoleSettingEquals
      static_val: '{{repl ConsoleSettingEquals "tls.key.name" "sample-value"}}'
    - name: ConsoleSettingNotEquals
      static_val: '{{repl ConsoleSettingNotEquals "tls.key.name" "sample-value"}}'
    - name: AppSetting
      static_val: '{{repl AppSetting "version.label"}}'

    # host start context
    - name: NodePublicIPAddressAll
      static_val: '{{repl NodePublicIPAddressAll "Elasticsearch" "elasticsearch"}}'
    - name: NodePublicIPAddressFirst
      static_val: '{{repl NodePublicIPAddressFirst "Elasticsearch" "elasticsearch"}}'
    - name: NodePublicIPAddress
      static_val: '{{repl NodePublicIPAddress "Elasticsearch" "elasticsearch"}}'
    - name: NodePrivateIPAddressAll
      static_val: '{{repl NodePrivateIPAddressAll "Elasticsearch" "elasticsearch"}}'
    - name: NodePrivateIPAddressFirst
      static_val: '{{repl NodePrivateIPAddressFirst "Elasticsearch" "elasticsearch"}}'
    - name: NodePrivateIPAddress
      static_val: '{{repl NodePrivateIPAddress "Elasticsearch" "elasticsearch"}}'
    - name: HostPublicIpAddress
      static_val: '{{repl HostPublicIpAddress "Elasticsearch" "elasticsearch"}}'
    - name: HostPublicIpAddressAll
      static_val: '{{repl HostPublicIpAddressAll "Elasticsearch" "elasticsearch"}}'
    - name: HostPrivateIpAddress
      static_val: '{{repl HostPrivateIpAddress "Elasticsearch" "elasticsearch"}}'
    - name: HostPrivateIpAddressAll
      static_val: '{{repl HostPrivateIpAddressAll "Elasticsearch" "elasticsearch"}}'
    - name: ContainerExposedPort
      static_val: '{{repl ContainerExposedPort "Elasticsearch" "elasticsearch" "9200"}}'
    - name: ContainerExposedPortAll
      static_val: '{{repl ContainerExposedPortAll "Elasticsearch" "elasticsearch" "9200"}}'
    - name: ContainerExposedPortFirst
      static_val: '{{repl ContainerExposedPortFirst "Elasticsearch" "elasticsearch" "9200"}}'

    # start context
    - name: ThisNodePublicIPAddress
      static_val: '{{repl ThisNodePublicIPAddress}}'
    - name: ThisNodeDockerAddress
      static_val: '{{repl ThisNodeDockerAddress}}'
    - name: ThisHostPublicIpAddress
      static_val: '{{repl ThisHostPublicIpAddress}}'
    - name: ThisHostPrivateIpAddress
      static_val: '{{repl ThisHostPrivateIpAddress}}'
    - name: HostPrivateIpAddressAll
      static_val: '{{repl ThisHostPrivateIpAddress}}'
    - name: HostPrivateIpAddressAll
      static_val: '{{repl ThisNodeInterfaceAddress "docker0"}}'
    - name: ThisHostInterfaceAddress
      static_val: '{{repl ThisHostInterfaceAddress "docker0"}}'
    - name: HostPrivateIpAddressAll
      static_val: '{{repl InterfaceAddress "docker0"}}'
`

	yamlValidator := validator.New(
		&validator.Config{TagName: "validate"},
	)

	var rootConfig libyaml.RootConfig
	if err := yaml.Unmarshal([]byte(s), &rootConfig); err != nil {
		t.Errorf("Unable to unmarshall yaml: ", err)
		return
	}

	if err := libyaml.RegisterValidations(yamlValidator); err != nil {
		t.Errorf("Failed to register validator: ", err)
		return
	}

	if ve, ok := yamlValidator.Struct(&rootConfig).(validator.ValidationErrors); ok {
		if len(ve) > 0 {
			t.Errorf("Found error during validation of good yaml: %s", ve)
		}
	}
}

func TestFailedValidation(t *testing.T) {
	s := `
replicated_api_version: "1.0.0"
name: validate_templ_funcs
components:
- name: Elasticsearch
  containers:
  - source: public
    image_name: elasticsearch
    version: 1.5.0-3
    ports:
    - public_port: 9200
      private_port: 9200
  - source: public
    image_name: plugin-head
    version: 0.2.0
    env_vars:
    # host start context
    - name: NodePublicIPAddressAll
      static_val: '{{repl NodePublicIPAddressAll "not-found" "elasticsearch"}}'
    - name: MakeBelieveFunction
      static_val: '{{repl MakeBelieveFunction}}'
    - name: ContainerExposedPort
      static_val: '{{repl ContainerExposedPort "Elasticsearch" "elasticsearch" "100000"}}'
`

	yamlValidator := validator.New(
		&validator.Config{TagName: "validate"},
	)

	var rootConfig libyaml.RootConfig
	if err := yaml.Unmarshal([]byte(s), &rootConfig); err != nil {
		t.Errorf("Unable to unmarshall yaml: ", err)
		return
	}

	if err := libyaml.RegisterValidations(yamlValidator); err != nil {
		t.Errorf("Failed to register validator: ", err)
		return
	}

	if ve, ok := yamlValidator.Struct(&rootConfig).(validator.ValidationErrors); ok {
		if !hasError(ve, "NodePublicIPAddressAll") {
			t.Error("Should have been an error for NodePublicIPAddressAll1")
		}
		if !hasError(ve, "MakeBelieveFunction") {
			t.Error("Should have been an error for MakeBelieveFunction as it doesnt exist")
		}
		if !hasError(ve, "ContainerExposedPort") {
			t.Error("Should have been an error for ContainerExposedPort as it refers to an invalid port number")
		}
	} else {
		t.Error("Failed to execute validaton check")
	}
}

func hasError(ve validator.ValidationErrors, match string) bool {
	for _, err := range ve {
		if strings.Contains(err.Value.(string), match) {
			return true
		}
	}
	return false
}
