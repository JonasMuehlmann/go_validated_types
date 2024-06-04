package main

type RangeNumber[T Number] struct {
	ValidatedNumber[T]
}

func NewRangeNumber[T Number](value T, min T, max T) RangeNumber[T] {
	object := RangeNumber[T]{}

	object.value = value
	if DEBUG {
		object.validator = MakeDefaultRangeValidator(min, max)
	}

	return object
}

func NewRangeNumberBare[T Number](min T, max T) RangeNumber[T] {
	object := RangeNumber[T]{}

	if DEBUG {
		object.validator = MakeDefaultRangeValidator(min, max)
	}

	return object
}
