package valid

type ValidatedType[T any] interface {
	// Validate checks if the value is valid
	Validate()
	// SetValidator sets the validator
	SetValidator(Validator[T])
	// GetValidator returns the validator
	GetValidator() Validator[T]
	// Get returns the value
	Get() T
	// Set sets the value and validates it
	Set(T)
}
