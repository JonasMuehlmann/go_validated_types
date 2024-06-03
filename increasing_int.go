package main

type IncreasingInt struct {
	ValidatedInt
}

func NewIncreasingInt(value int) IncreasingInt {
	object := IncreasingInt{}

	object.value = value

	return object
}

func NewIncreasingIntBare() IncreasingInt {
	object := IncreasingInt{}

	return object
}

func (v *IncreasingInt) Set(value int) {
	if DEBUG {
		if value < v.value {
			panic("Value must be increasing")
		}
	}

	v.ValidatedInt.Set(value)
}
