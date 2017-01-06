package libyaml_test

import (
	"testing"

	validator "gopkg.in/go-playground/validator.v8"
)

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
