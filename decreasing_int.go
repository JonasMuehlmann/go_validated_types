package main

type DecreasingInt struct {
	ValidatedInt
}

func NewDecreasingInt(value int) DecreasingInt {
	object := DecreasingInt{}

	object.value = value

	return object
}

func NewDecreasingIntBare() DecreasingInt {
	object := DecreasingInt{}

	return object
}

func (v *DecreasingInt) Set(value int) {
	if DEBUG {
		if value > v.value {
			panic("Value must be decreasing")
		}
	}

	v.ValidatedInt.Set(value)
}
