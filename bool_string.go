package libyaml

import "strconv"

type BoolString string

func (s BoolString) ParseBool() (bool, error) {
	return strconv.ParseBool(string(s))
}

func (s BoolString) MarshalYAML() (interface{}, error) {
	if s == "" {
		return false, nil
	}
	b, err := s.ParseBool()
	if err == nil {
		return b, nil
	}
	return s, nil
}
