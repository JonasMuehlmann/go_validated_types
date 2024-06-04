package main

type TrimmedString struct {
	ValidatedString
}

func NewTrimmedString(value string) TrimmedString {
	object := TrimmedString{}
	object.value = value

	if DEBUG {
		object.validator = MakeTrimmedValidator()
		object.validate()
	}

	return object
}

func NewTrimmedStringBare() TrimmedString {
	object := TrimmedString{}

	if DEBUG {
		object.validator = MakeTrimmedValidator()
		object.validate()
	}

	return object
}
