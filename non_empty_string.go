package valid

type NonEmptyString struct {
	ValidatedString
}

func NewNonEmptyString(value string, prefix string) NonEmptyString {
	object := NonEmptyString{}
	object.value = value

	if DEBUG {
		object.validator = MakeNonEmptyStringValidator()
		object.validate()
	}

	return object
}
