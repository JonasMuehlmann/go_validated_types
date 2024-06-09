package valid

type SetOnce[T any] struct {
	ValidatedTypeBase[T]
}

func NewSetOnce[T any](value T) SetOnce[T] {
	object := SetOnce[T]{}

	object.Set(value)

	return object
}

func NewSetOnceBare[T any]() SetOnce[T] {
	return SetOnce[T]{}
}

func (v *SetOnce[T]) Set(value T) {
	v.ValidatedTypeBase.Set(value)
	if DEBUG {
		v.SetValidator(MakeContradictionValidator[T]("Value can only be set once"))
	}
}

func (v *SetOnce[T]) Get() T {
	if DEBUG {
		if v.GetValidator() == nil {
			panic("Value has not been set")
		}
	}

	return v.ValidatedTypeBase.Get()
}
