package main

type DecreasingNumber[T Number] struct {
	ValidatedNumber[T]
}

func NewDecreasingNumber[T Number](value T) DecreasingNumber[T] {
	object := DecreasingNumber[T]{}

	object.value = value

	return object
}

func NewDecreasingNumberBare[T Number]() DecreasingNumber[T] {
	object := DecreasingNumber[T]{}

	return object
}

func (v *DecreasingNumber[T]) Set(value T) {
	if DEBUG {
		if value > v.value {
			panic("Value must be decreasing")
		}
	}

	v.ValidatedNumber.Set(value)
}
