package main

type RangeInt struct {
	ValidatedInt
}

func NewRangeInt(value int, min int, max int) RangeInt {
	object := RangeInt{}

	object.value = value
	object.validator = MakeDefaultRangeValidator(min, max)

	return object
}

func NewRangeIntBare(min int, max int) RangeInt {
	object := RangeInt{}

	object.validator = MakeDefaultRangeValidator(min, max)

	return object
}
