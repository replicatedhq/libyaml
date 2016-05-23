package libyaml

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/blang/semver"
	validator "gopkg.in/bluesuncorp/validator.v8"
)

var (
	keyRe   = regexp.MustCompile("^([^\\[]+)(?:\\[(\\d+)\\])?$")
	bytesRe = regexp.MustCompile(`(?i)^(-?\d+)([KMGT]B?|B)$`)
)

// RegisterValidations will register all known validation for the libyaml project.
func RegisterValidations(v *validator.Validate) error {
	if err := v.RegisterValidation("configitemtype", ConfigItemTypeValidation); err != nil {
		return err
	}
	if err := v.RegisterValidation("configitemwhen", ConfigItemWhenValidation); err != nil {
		return err
	}

	if err := v.RegisterValidation("apiversion", SemverValidation); err != nil {
		return err
	}

	if err := v.RegisterValidation("semver", SemverValidation); err != nil {
		return err
	}

	if err := v.RegisterValidation("componentexists", ComponentExistsValidation); err != nil {
		return err
	}

	if err := v.RegisterValidation("containerexists", ContainerExistsValidation); err != nil {
		return err
	}

	if err := v.RegisterValidation("componentcontainer", ComponentContainerFormatValidation); err != nil {
		return err
	}

	if err := v.RegisterValidation("absolutepath", IsAbsolutePathValidation); err != nil {
		return err
	}

	// will handle this in vendor web. this prevents panic from validator.v8 library
	if err := v.RegisterValidation("integrationexists", NoopValidation); err != nil {
		return err
	}

	// will handle this in vendor web. this prevents panic from validator.v8 library
	if err := v.RegisterValidation("externalregistryexists", NoopValidation); err != nil {
		return err
	}

	if err := v.RegisterValidation("bytes", IsBytesValidation); err != nil {
		return err
	}

	if err := v.RegisterValidation("bool", IsBoolValidation); err != nil {
		return err
	}

	if err := v.RegisterValidation("tcpport", IsTCPPortValidation); err != nil {
		return err
	}

	if err := v.RegisterValidation("graphiteretention", GraphiteRetentionFormatValidation); err != nil {
		return err
	}

	if err := v.RegisterValidation("graphiteaggregation", GraphiteAggregationFormatValidation); err != nil {
		return err
	}

	if err := v.RegisterValidation("monitorlabelscale", MonitorLabelScaleValidation); err != nil {
		return err
	}

	return nil
}

func FormatFieldError(key string, fieldErr *validator.FieldError, root *RootConfig) error {
	formatted, err := FormatKey(key, fieldErr, root)
	if err != nil {
		formatted = key
	}

	switch fieldErr.Tag {
	case "apiversion":
		return fmt.Errorf("A valid \"replicated_api_version\" is required as a root element")

	case "componentexists":
		return fmt.Errorf("Component \"%s\" does not exist at key \"%s\"", fieldErr.Value, formatted)

	case "containerexists":
		return fmt.Errorf("Container \"%s\" does not exist at key \"%s\"", fieldErr.Value, formatted)

	case "componentcontainer":
		return fmt.Errorf("Should be in the format \"<component name>,<container image name>\" at key \"%s\"", formatted)

	case "integrationexists":
		return fmt.Errorf("Missing integration \"%s\" at key \"%s\"", fieldErr.Value, formatted)

	case "externalregistryexists":
		return fmt.Errorf("Missing external registry integration \"%s\" at key \"%s\"", fieldErr.Value, formatted)

	case "required":
		return fmt.Errorf("Value required at key \"%s\"", formatted)

	case "tcpport":
		return fmt.Errorf("A valid port number must be between 0 and 65535: %q", formatted)

	case "graphiteretention":
		return fmt.Errorf("Should be in the new style graphite retention policy at key %q", formatted)

	case "graphiteaggregation":
		return fmt.Errorf("Valid values for graphite aggregation method are 'average', 'sum', 'min', 'max', 'last' at key %q", formatted)

	case "monitorlabelscale":
		return fmt.Errorf("Please specify 'metric', 'none', or a floating point number for scale at %q", formatted)

	default:
		return fmt.Errorf("Validation failed on the \"%s\" tag at key \"%s\"", fieldErr.Tag, formatted)
	}
}

func FormatKey(keyChain string, fieldErr *validator.FieldError, root *RootConfig) (string, error) {
	value := reflect.ValueOf(*root)
	keys := strings.Split(keyChain, ".")

	rest, err := formatKey(keys, value)
	if err != nil {
		return "", err
	}

	if rest != "" {
		rest = rest[1:]

		matches := keyRe.FindStringSubmatch(fieldErr.Field)
		if matches[2] != "" {
			rest += fmt.Sprintf("[%s]", matches[2])
		}
	}

	return rest, nil
}

func formatKey(keys []string, parent reflect.Value) (string, error) {
	if len(keys) == 1 {
		return "", nil
	}

	if parent.Type().Kind() == reflect.Ptr {
		parent = parent.Elem()
	}

	if parent.Type().Kind() == reflect.Struct {
		key := keys[1]
		matches := keyRe.FindStringSubmatch(key)

		field, ok := parent.Type().FieldByName(matches[1])
		if !ok {
			return "", fmt.Errorf("field \"%s\" not found", matches[1])
		}

		yamlTag := field.Tag.Get("yaml")

		value := parent.FieldByName(matches[1])

		rest, err := formatKey(keys[1:], value)
		if err != nil {
			return "", err
		}

		return "." + yamlTag + rest, nil
	} else if parent.Type().Kind() == reflect.Slice {
		key := keys[0]
		matches := keyRe.FindStringSubmatch(key)

		i, err := strconv.Atoi(matches[2])
		if err != nil {
			return "", err
		}

		value := parent.Index(i)

		rest, err := formatKey(keys, value)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("[%d]", i) + rest, nil
	}

	return "", nil
}

// ConfigItemTypeValidation will validate that the type element of a config item is a supported and valid option.
func ConfigItemTypeValidation(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if fieldKind != reflect.String {
		return false
	}

	validTypes := map[string]bool{
		"text":        true,
		"label":       true,
		"password":    true,
		"file":        true,
		"bool":        true,
		"select_one":  true,
		"select_many": true,
		"textarea":    true,
		"select":      true,
	}

	if validTypes[field.String()] {
		return true
	}

	return false
}

// ConfigItemWhenValidation will validate that the when element of a config item is in a valid format and references other valid, created objects.
func ConfigItemWhenValidation(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	root, ok := topStruct.Interface().(*RootConfig)
	if !ok {
		// this is an issue with the code and really should be a panic
		return true
	}

	if fieldKind != reflect.String {
		// this is an issue with the code and really should be a panic
		return true
	}

	var whenValue string

	whenValue = field.String()
	if whenValue == "" {
		return true
	}

	splitString := "="
	if strings.Contains(whenValue, "!=") {
		splitString = "!="
	}

	parts := strings.SplitN(whenValue, splitString, 2)
	if len(parts) >= 2 {
		whenValue = parts[0]
	}

	return configItemExists(whenValue, root)
}

func configItemExists(configItemName string, root *RootConfig) bool {
	for _, group := range root.ConfigGroups {
		for _, item := range group.Items {
			if item != nil && item.Name == configItemName {
				return true
			}
			if item != nil {
				for _, childItem := range item.Items {
					if childItem != nil && childItem.Name == configItemName {
						return true
					}
				}
			}
		}
	}

	return false
}

func componentExists(componentName string, root *RootConfig) bool {
	for _, component := range root.Components {
		if component != nil && component.Name == componentName {
			return true
		}
	}

	return false
}

func containerExists(componentName, containerName string, root *RootConfig) bool {
	for _, component := range root.Components {
		if component != nil && component.Name == componentName {
			for _, container := range component.Containers {
				if container != nil && container.ImageName == containerName {
					return true
				}
			}
			return false
		}
	}

	return false
}

// ComponentExistsValidation will validate that the specified component name is present in the current YAML.
func ComponentExistsValidation(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	// validates that the component exists in the root.Components slice
	root, ok := topStruct.Interface().(*RootConfig)
	if !ok {
		// this is an issue with the code and really should be a panic
		return true
	}

	if fieldKind != reflect.String {
		// this is an issue with the code and really should be a panic
		return true
	}

	var componentName string

	componentName = field.String()

	parts := strings.SplitN(componentName, ",", 2)

	if len(parts) >= 2 {
		componentName = parts[0]
	}

	return componentExists(componentName, root)
}

// ContainerExistsValidation will validate that the specified container name is present in the current YAML.
func ContainerExistsValidation(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	// validates that the container exists in the root.components.containers slice

	root, ok := topStruct.Interface().(*RootConfig)
	if !ok {
		// this is an issue with the code and really should be a panic
		return true
	}

	var componentName, containerName string

	if fieldKind != reflect.String {
		// this is an issue with the code and really should be a panic
		return true
	}

	containerName = field.String()

	if param != "" {
		componentField, componentKind, ok := v.GetStructFieldOK(currentStructOrField, param)

		if !ok || componentKind != reflect.String {
			// this is an issue with the code and really should be a panic
			return true
		}

		componentName = componentField.String()
	} else {
		parts := strings.SplitN(containerName, ",", 2)

		if len(parts) < 2 {
			// let "componentcontainer" validation handle this case
			return true
		}

		componentName = parts[0]
		containerName = parts[1]
	}

	if !componentExists(componentName, root) {
		// let "componentexists" validation handle this case
		return true
	}

	return containerExists(componentName, containerName, root)
}

// IsAbsolutePathValidation validates that the format of the field begins with a "/"
func IsAbsolutePathValidation(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if fieldKind != reflect.String {
		// this is an issue with the code and really should be a panic
		return true
	}

	return strings.HasPrefix(field.String(), "/")
}

// ComponentContainerFormatValidation will validate that component/container name is in the correct format.
func ComponentContainerFormatValidation(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	// validates the format of the string field conforms to "<component name>,<container image name>"

	if fieldKind != reflect.String {
		// this is an issue with the code and really should be a panic
		return true
	}

	parts := strings.SplitN(field.String(), ",", 2)

	if len(parts) < 2 {
		return false
	}

	return true
}

// SemverValidation will validate that the field is in correct, proper semver format.
func SemverValidation(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if fieldKind != reflect.String {
		return true
	}

	if field.String() == "" {
		return true
	}

	_, err := semver.Make(field.String())
	return err == nil
}

// IsBytesValidation will return if a field is a parseable bytes value.
func IsBytesValidation(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if fieldKind != reflect.String {
		// this is an issue with the code and really should be a panic
		return true
	}

	if field.String() == "" {
		return true
	}

	parts := bytesRe.FindStringSubmatch(strings.TrimSpace(field.String()))
	if len(parts) < 3 {
		return false
	}

	value, err := strconv.ParseUint(parts[1], 10, 0)
	if err != nil || value < 1 {
		return false
	}

	return true
}

// IsBoolValidation will return if a string field parses to a bool.
func IsBoolValidation(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if fieldKind != reflect.String {
		// this is an issue with the code and really should be a panic
		return true
	}

	if field.String() == "" {
		return true
	}

	_, err := strconv.ParseBool(field.String())
	if err != nil {
		return false
	}

	return true
}

// IsTCPPortValidation will return true if a field value is also a valid TCP port.
func IsTCPPortValidation(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if fieldKind != reflect.Int32 {
		// this is an issue with the code and really should be a panic
		return true
	}

	port := field.Int()
	return 0 <= port && port <= 65535
}

// NoopValidation will return true always.
func NoopValidation(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	return true
}

// GraphiteRetentionFormatValidation will return true if the field value is a valid graphite retention value.
func GraphiteRetentionFormatValidation(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if fieldKind != reflect.String {
		return true
	}

	valueStr := field.String()
	if valueStr == "" {
		return true
	}

	validate := func(amount, unit string) bool {
		_, err := strconv.Atoi(amount)
		if err != nil {
			return false
		}

		switch unit {
		case "s", "m", "h", "d", "y":
			return true
		default:
			return false
		}
	}

	// Example: 15s:7d,1m:21d,15m:5y
	periods := strings.Split(valueStr, ",")
	for _, period := range periods {
		periodParts := strings.Split(period, ":")
		if len(periodParts) != 2 {
			return false
		}

		for _, part := range periodParts {
			partLen := len(part)
			if partLen < 2 {
				return false
			}
			if !validate(part[:partLen-1], part[partLen-1:]) {
				return false
			}
		}
	}

	return true
}

// GraphiteAggregationFormatValidation will return true if the field value is a valid value for the a graphite aggregation.
func GraphiteAggregationFormatValidation(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if fieldKind != reflect.String {
		return true
	}

	valueStr := field.String()
	if valueStr == "" {
		return true
	}

	switch valueStr {
	case "average", "sum", "min", "max", "last":
		return true
	default:
		return false
	}
}

// MonitorLabelScaleValidation will return true only if the value is a parseable and correct value for the scale.
func MonitorLabelScaleValidation(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if fieldKind != reflect.String {
		return true
	}

	valueStr := field.String()
	if valueStr == "" {
		return true
	}

	switch valueStr {
	case "metric", "none":
		return true
	default:
		_, err := strconv.ParseFloat(valueStr, 64)
		return err == nil
	}
}
