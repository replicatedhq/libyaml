package libyaml_test

import (
	"fmt"
	"testing"

	. "github.com/replicatedhq/libyaml"
	validator "gopkg.in/go-playground/validator.v8"
)

func TestIntValidation(t *testing.T) {
	v := validator.New(&validator.Config{TagName: "validate"})
	err := v.RegisterValidation("int", IntValidation)
	if err != nil {
		t.Fatal(err)
	}
	err = v.Field("123", "int")
	if err != nil {
		t.Errorf("got unexpected error %v", err)
	}
	err = v.Field("-123", "int")
	if err != nil {
		t.Errorf("got unexpected error %v", err)
	}
	err = v.Field("", "int")
	if err := AssertValidationErrors(t, err, map[string]string{
		"": "int",
	}); err != nil {
		t.Error(err)
	}
	err = v.Field("", "omitempty,int")
	if err != nil {
		t.Errorf("got unexpected error %v", err)
	}
	err = v.Field("123.1", "int")
	if err := AssertValidationErrors(t, err, map[string]string{
		"": "int",
	}); err != nil {
		t.Error(err)
	}
	err = v.Field("abc", "int")
	if err := AssertValidationErrors(t, err, map[string]string{
		"": "int",
	}); err != nil {
		t.Error(err)
	}
}

func TestDockerVersionValidation(t *testing.T) {
	v := validator.New(&validator.Config{TagName: "validate"})
	err := v.RegisterValidation("dockerversion", DockerVersionValidation)
	if err != nil {
		t.Fatal(err)
	}
	err = v.Field("1.13.1", "dockerversion")
	if err != nil {
		t.Errorf("got unexpected error %v", err)
	}
	err = v.Field("1.0.0", "dockerversion")
	if err != nil {
		t.Errorf("got unexpected error %v", err)
	}
	err = v.Field("17.03.0-ce", "dockerversion")
	if err != nil {
		t.Errorf("got unexpected error %v", err)
	}
	err = v.Field("17.03.0-ee", "dockerversion")
	if err != nil {
		t.Errorf("got unexpected error %v", err)
	}
	err = v.Field("17.03.0", "dockerversion") // seems important not to have to specify edition here
	if err != nil {
		t.Errorf("got unexpected error %v", err)
	}
	err = v.Field("", "omitempty,dockerversion")
	if err != nil {
		t.Errorf("got unexpected error %v", err)
	}
	err = v.Field("", "dockerversion")
	if err := AssertValidationErrors(t, err, map[string]string{
		"": "dockerversion",
	}); err != nil {
		t.Error(err)
	}
	err = v.Field("blah", "dockerversion")
	if err := AssertValidationErrors(t, err, map[string]string{
		"": "dockerversion",
	}); err != nil {
		t.Error(err)
	}
	err = v.Field("17.13.1-ce", "dockerversion")
	if err := AssertValidationErrors(t, err, map[string]string{
		"": "dockerversion",
	}); err != nil {
		t.Error(err)
	}
	err = v.Field("0.1.1", "dockerversion")
	if err := AssertValidationErrors(t, err, map[string]string{
		"": "dockerversion",
	}); err != nil {
		t.Error(err)
	}
	err = v.Field("1.14.1", "dockerversion")
	if err := AssertValidationErrors(t, err, map[string]string{
		"": "dockerversion",
	}); err != nil {
		t.Error(err)
	}
	err = v.Field("1.13.1-alpha", "dockerversion") // idk about pinning weird modified versions
	if err := AssertValidationErrors(t, err, map[string]string{
		"": "dockerversion",
	}); err != nil {
		t.Error(err)
	}
}

type RequiredMinAPIVersionStruct struct {
	APIVersion string
	Required   string `validate:"required_minapiversion=2.8.0"`
}

func (r *RequiredMinAPIVersionStruct) GetAPIVersion() string {
	return r.APIVersion
}

func TestRequiredMinAPIVersionValidation(t *testing.T) {
	v := validator.New(&validator.Config{TagName: "validate"})
	err := v.RegisterValidation("required_minapiversion", RequiredMinAPIVersion)
	if err != nil {
		t.Fatal(err)
	}
	err = v.Struct(&RequiredMinAPIVersionStruct{APIVersion: "2.8.0"})
	if err != nil {
		t.Errorf("got unexpected error %v", err)
	}
	err = v.Struct(&RequiredMinAPIVersionStruct{APIVersion: "2.8.1"})
	if err != nil {
		t.Errorf("got unexpected error %v", err)
	}
	err = v.Struct(&RequiredMinAPIVersionStruct{APIVersion: "2.7.0"})
	if err := AssertValidationErrors(t, err, map[string]string{
		"RequiredMinAPIVersionStruct.Required": "required_minapiversion",
	}); err != nil {
		t.Error(err)
	}
}

func AssertValidationErrors(t *testing.T, err error, pathAndTags map[string]string) error {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return fmt.Errorf("expecting validator.ValidationErrors, got %T", err)
	}
	var multiErr MultiError
	if len(validationErrors) != len(pathAndTags) {
		multiErr.Append(fmt.Errorf("expecting validator.ValidationErrors length %d, got %d", len(pathAndTags), len(validationErrors)))
	}
	for path, tag := range pathAndTags {
		err, ok := validationErrors[path]
		if !ok {
			multiErr.Append(fmt.Errorf("validator.ValidationErrors at path %s not found", path))
			continue
		}
		if err.Tag != tag {
			multiErr.Append(fmt.Errorf("expecting validator.ValidationErrors at path %s to have tag %s, got tag %s", path, tag, err.Tag))
		}
	}
	return multiErr.ErrorOrNil()
}
