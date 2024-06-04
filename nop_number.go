package main

type NopNumber[T Number] struct {
	ValidatedNumber[T]
}

func NewNopNumber[T Number](value T) NopNumber[T] {
	object := NopNumber[T]{}

	object.value = value

	return object
}

func NewNopNumberBare[T Number]() NopNumber[T] {
	return NopNumber[T]{}
}
