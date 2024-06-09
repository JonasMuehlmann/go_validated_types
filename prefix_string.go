package valid

type PrefixString struct {
	ValidatedString
}

func NewPrefixString(value string, prefix string) PrefixString {
	object := PrefixString{}
	object.value = value

	if DEBUG {
		object.validator = MakePrefixValidator(prefix)
		object.validate()
	}

	return object
}
