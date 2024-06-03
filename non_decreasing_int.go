package main

type NonDecreasingInt struct {
	ValidatedInt
}

func NewNonDecreasingInt(value int) NonDecreasingInt {
	object := NonDecreasingInt{}

	object.value = value

	return object
}

func NewNonDecreasingIntBare() NonDecreasingInt {
	object := NonDecreasingInt{}

	return object
}

func (v *NonDecreasingInt) Set(value int) {
	if DEBUG {
		if value >= v.value {
			panic("Value must not be nondecreasing")
		}
	}

	v.ValidatedInt.Set(value)
}
