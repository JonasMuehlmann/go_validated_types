package main

type ValidatedInt struct {
	ValidatedType[int]
	value     int
	validator Validator
}

func (v *ValidatedInt) validate() {
	if v.validator != nil {
		v.validator(v.value)
	}
}

func (v *ValidatedInt) Get() int {
	return v.value
}

func (v *ValidatedInt) Set(value int) {
	v.value = value
	if DEBUG {
		v.validate()
	}
}

func NewValidatedInt(value int, validator Validator) ValidatedInt {
	object := ValidatedInt{value: value}

	if DEBUG {
		object.validator = validator
		if object.validator != nil {
			object.validate()
		}
	}

	return object
}
