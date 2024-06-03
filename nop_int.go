package main

type NopInt struct {
	ValidatedInt
}

func NewNopInt(value int) NopInt {
	return NopInt{NewValidatedInt(value, MakeTautologyValidator())}
}
