package valid

type NonIncreasingNumber[T Number] struct {
	ValidatedNumber[T]
}

func NewNonIncreasingNumber[T Number](value T) NonIncreasingNumber[T] {
	object := NonIncreasingNumber[T]{}

	object.value = value

	return object
}

func NewNonIncreasingNumberBare[T Number]() NonIncreasingNumber[T] {
	object := NonIncreasingNumber[T]{}

	return object
}

func (v *NonIncreasingNumber[T]) Set(value T) {
	if DEBUG {
		if value <= v.value {
			panic("Value must be nonincreasing")
		}
	}

	v.ValidatedNumber.Set(value)
}
