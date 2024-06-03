package main

type Validator func(int)

// MakeRangeValidator returns a validator that checks if a value is within a range.
func MakeRangeValidator(min int, max int, message string) Validator {
	return func(value int) {
		if !(value >= min && value <= max) {
			panic(message)
		}
	}
}

// MakeDefaultRangeValidator returns a validator that checks if a value is within a range.
func MakeDefaultRangeValidator(min int, max int) Validator {
	return MakeRangeValidator(min, max, "Value is out of range")
}

// MakeTautologyValidator returns a validator that always passes.
func MakeTautologyValidator() Validator {
	return func(value int) {}
}

// MakeContradictionValidator returns a validator that always fails.
func MakeContradictionValidator(message string) Validator {
	return func(value int) {
		panic(message)
	}
}

// MakeDefaultContradictionValidator returns a validator that always fails.
func MakeDefaultContradictionValidator() Validator {
	return MakeContradictionValidator("Validation is set to always fail")
}

// MakeNilValidator returns a nil validator.
func MakeNilValidator() Validator {
	return nil
}
