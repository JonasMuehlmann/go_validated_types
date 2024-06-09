package valid

type SuffixString struct {
	ValidatedString
}

func NewSuffixString(value string, suffix string) SuffixString {
	object := SuffixString{}
	object.value = value

	if DEBUG {
		object.validator = MakeSuffixValidator(suffix)
		object.validate()
	}

	return object
}
