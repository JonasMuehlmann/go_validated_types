package main

type ValidatedTypeBase[T any] struct {
	value     T
	validator Validator[T]
}

func (v ValidatedTypeBase[T]) Validate() {
	if DEBUG {
		if v.validator != nil {
			v.validator(v.value)
		}
	}
}

func (v *ValidatedTypeBase[T]) SetValidator(validator Validator[T]) {
	if DEBUG {
		v.validator = validator
	}
}

func (v ValidatedTypeBase[T]) GetValidator() Validator[T] {
	return v.validator
}

func (v ValidatedTypeBase[T]) Get() T {
	return v.value
}

func (v *ValidatedTypeBase[T]) Set(t T) {
	v.value = t
	if DEBUG {
		v.Validate()
	}
}

func NewValidatedTypeBase[T any](value T) ValidatedTypeBase[T] {
	object := ValidatedTypeBase[T]{}

	object.Set(value)

	if DEBUG {
		object.SetValidator(MakeContradictionValidator[T]("Value can only be set once"))
	}

	return object
}

func NewValidatedTypeBaseBare[T any]() ValidatedTypeBase[T] {
	return ValidatedTypeBase[T]{}
}
