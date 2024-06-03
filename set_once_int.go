package main

type SetOnceInt struct {
	ValidatedInt
}

func NewSetOnceInt(value int) SetOnceInt {
	object := SetOnceInt{}

	object.value = value

	return object
}

func NewSetOnceIntBare() SetOnceInt {
	return SetOnceInt{}
}

func (v *SetOnceInt) Set(value int) {
	v.ValidatedInt.Set(value)
	if DEBUG {
		v.validator = MakeContradictionValidator("Value can only be set once")
	}
}

func (v *SetOnceInt) Get() int {
	if DEBUG {
		if v.validator == nil {
			panic("Value has not been set")
		}
	}

	return v.value
}
