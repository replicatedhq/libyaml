package libyaml_test

import (
	"fmt"
	"testing"

	. "github.com/replicatedhq/libyaml"
	validator "gopkg.in/go-playground/validator.v8"
)

func TestIntValidator(t *testing.T) {
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
