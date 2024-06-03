package main

type ValidatedType[T any] interface {
	// Validate checks if the value is valid
	Validate()
	// Get returns the value
	Get() T
	// Set sets the value and validates it
	Set(T)
}
