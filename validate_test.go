package libyaml_test

import (
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
	AssertValidationErrors(t, err, map[string]string{
		"": "int",
	})
	err = v.Field("", "omitempty,int")
	if err != nil {
		t.Errorf("got unexpected error %v", err)
	}
	err = v.Field("123.1", "int")
	AssertValidationErrors(t, err, map[string]string{
		"": "int",
	})
	err = v.Field("abc", "int")
	AssertValidationErrors(t, err, map[string]string{
		"": "int",
	})
}

func AssertValidationErrors(t *testing.T, err error, pathAndTags map[string]string) bool {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		t.Errorf("expecting validator.ValidationErrors, got %T", err)
		return false
	}
	passed := true
	if len(validationErrors) != len(pathAndTags) {
		t.Errorf("expecting validator.ValidationErrors length %d, got %d", len(pathAndTags), len(validationErrors))
		passed = false
	}
	for path, tag := range pathAndTags {
		err, ok := validationErrors[path]
		if !ok {
			t.Errorf("validator.ValidationErrors at path %s not found", path)
			passed = false
			continue
		}
		if err.Tag != tag {
			t.Errorf("expecting validator.ValidationErrors at path %s to have tag %s, got tag %s", path, tag, err.Tag)
			passed = false
		}
	}
	return passed
}
