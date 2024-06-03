package main

type RangeInt struct {
	ValidatedInt
}

func NewRangeInt(value int, min int, max int) RangeInt {
	object := RangeInt{}

	object.value = value
	if DEBUG {
		object.validator = MakeDefaultRangeValidator(min, max)
	}

	return object
}

func NewRangeIntBare(min int, max int) RangeInt {
	object := RangeInt{}

	if DEBUG {
		object.validator = MakeDefaultRangeValidator(min, max)
	}

	return object
}
