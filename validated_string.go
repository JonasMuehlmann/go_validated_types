package valid

type ValidatedString struct {
	ValidatedType[string]
	value     string
	validator Validator[string]
}

func (v *ValidatedString) validate() {
	if v.validator != nil {
		v.validator(v.value)
	}
}

func (v *ValidatedString) Get() string {
	return v.value
}

func (v *ValidatedString) Set(value string) {
	v.value = value
	if DEBUG {
		v.validate()
	}
}

func NewValidatedString(value string, validator Validator[string]) ValidatedString {
	object := ValidatedString{value: value}

	if DEBUG {
		object.validator = validator
		if object.validator != nil {
			object.validate()
		}
	}

	return object
}

func NewValidatedStringBare(validator Validator[string]) ValidatedString {
	object := ValidatedString{}

	if DEBUG {
		object.validator = validator
		if object.validator != nil {
			object.validate()
		}
	}

	return object
}
