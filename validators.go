package main

import "golang.org/x/exp/constraints"

type Validator[T any] func(T)

// MakeRangeValidator returns a validator that checks if a value is within a range.
func MakeRangeValidator[T constraints.Ordered](min T, max T, message string) Validator[T] {
	return func(value T) {
		if !(value >= min && value <= max) {
			panic(message)
		}
	}
}

// MakeDefaultRangeValidator returns a validator that checks if a value is within a range.
func MakeDefaultRangeValidator[T constraints.Ordered](min T, max T) Validator[T] {
	return MakeRangeValidator(min, max, "Value is out of range")
}

// MakeTautologyValidator returns a validator that always passes.
func MakeTautologyValidator[T any]() Validator[T] {
	return func(value T) {}
}

// MakeContradictionValidator returns a validator that always fails.
func MakeContradictionValidator[T any](message string) Validator[T] {
	return func(value T) {
		panic(message)
	}
}

// MakeDefaultContradictionValidator returns a validator that always fails.
func MakeDefaultContradictionValidator[T any]() Validator[T] {
	return MakeContradictionValidator[T]("Validation is set to always fail")
}

// MakeNilValidator returns a nil validator.
func MakeNilValidator[T any]() Validator[T] {
	return nil
}
