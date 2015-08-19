package libyaml

type minInt1 uint

func (m *minInt1) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var i uint
	if err := unmarshal(&i); err != nil {
		return err
	}

	if i == 0 {
		*m = minInt1(1)
	} else {
		*m = minInt1(i)
	}

	return nil
}

func (m minInt1) MarshalYAML() (interface{}, error) {
	if m == 0 {
		return 1, nil
	}

	return m, nil
}
