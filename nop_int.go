package main

type NopInt struct {
	ValidatedInt
}

func NewNopInt(value int) NopInt {
	object := NopInt{}

	object.value = value

	return object
}

func NewNopIntBare() NopInt {
	return NopInt{}
}
