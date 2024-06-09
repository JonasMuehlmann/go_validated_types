package valid

type NonDecreasingNumber[T Number] struct {
	ValidatedNumber[T]
}

func NewNonDecreasingNumber[T Number](value T) NonDecreasingNumber[T] {
	object := NonDecreasingNumber[T]{}

	object.value = value

	return object
}

func NewNonDecreasingNumberBare[T Number]() NonDecreasingNumber[T] {
	object := NonDecreasingNumber[T]{}

	return object
}

func (v *NonDecreasingNumber[T]) Set(value T) {
	if DEBUG {
		if value >= v.value {
			panic("Value must not be nondecreasing")
		}
	}

	v.ValidatedNumber.Set(value)
}
