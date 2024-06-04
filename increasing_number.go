package main

type IncreasingNumber[T Number] struct {
	ValidatedNumber[T]
}

func NewIncreasingNumber[T Number](value T) IncreasingNumber[T] {
	object := IncreasingNumber[T]{}

	object.value = value

	return object
}

func NewIncreasingNumberBare[T Number]() IncreasingNumber[T] {
	object := IncreasingNumber[T]{}

	return object
}

func (v *IncreasingNumber[T]) Set(value T) {
	if DEBUG {
		if value < v.value {
			panic("Value must be increasing")
		}
	}

	v.ValidatedNumber.Set(value)
}
