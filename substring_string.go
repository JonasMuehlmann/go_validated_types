package valid

type SubstringString struct {
	ValidatedString
}

func NewSubstringString(value string, substring string) SubstringString {
	object := SubstringString{}
	object.value = value

	if DEBUG {
		object.validator = MakeSubstringValidator(substring)
		object.validate()
	}

	return object
}
