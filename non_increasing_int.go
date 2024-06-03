package main

type NonIncreasingInt struct {
	ValidatedInt
}

func NewNonIncreasingInt(value int) NonIncreasingInt {
	object := NonIncreasingInt{}

	object.value = value

	return object
}

func NewNonIncreasingIntBare() NonIncreasingInt {
	object := NonIncreasingInt{}

	return object
}

func (v *NonIncreasingInt) Set(value int) {
	if DEBUG {
		if value <= v.value {
			panic("Value must be nonincreasing")
		}
	}

	v.ValidatedInt.Set(value)
}
