package main

type SetOnceNumber[T Number] struct {
	ValidatedNumber[T]
}

func NewSetOnceNumber[T Number](value T) SetOnceNumber[T] {
	object := SetOnceNumber[T]{}

	object.value = value

	if DEBUG {
		object.validator = MakeContradictionValidator[T]("Value can only be set once")
	}

	return object
}

func NewSetOnceNumberBare[T Number]() SetOnceNumber[T] {
	return SetOnceNumber[T]{}
}

func (v *SetOnceNumber[T]) Set(value T) {
	v.ValidatedNumber.Set(value)
	if DEBUG {
		v.validator = MakeContradictionValidator[T]("Value can only be set once")
	}
}

func (v *SetOnceNumber[T]) Get() T {
	if DEBUG {
		if v.validator == nil {
			panic("Value has not been set")
		}
	}

	return v.value
}
