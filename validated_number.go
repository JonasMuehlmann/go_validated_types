package valid

type ValidatedNumber[T Number] struct {
	ValidatedType[T]
	value     T
	validator Validator[T]
}

func (v *ValidatedNumber[T]) validate() {
	if v.validator != nil {
		v.validator(v.value)
	}
}

func (v *ValidatedNumber[T]) Get() T {
	return v.value
}

func (v *ValidatedNumber[T]) Set(value T) {
	v.value = value
	if DEBUG {
		v.validate()
	}
}

func NewValidatedNumber[T Number](value T, validator Validator[T]) ValidatedNumber[T] {
	object := ValidatedNumber[T]{value: value}

	if DEBUG {
		object.validator = validator
		if object.validator != nil {
			object.validate()
		}
	}

	return object
}
func NewValidatedNumberBare[T Number](validator Validator[T]) ValidatedNumber[T] {
	object := ValidatedNumber[T]{}

	if DEBUG {
		object.validator = validator
		if object.validator != nil {
			object.validate()
		}
	}

	return object
}
