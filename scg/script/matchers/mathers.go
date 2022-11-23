package matchers

var linkNameWithMatherName = map[string]string{
	"number":                   "NumberMatcher",
	"positive_number":          "PositiveNumberMatcher",
	"positive_number_in_words": "PositiveNumberInWordsMatcher",
	"any":                      "AnyMatcher",
	"first_word":               "FirstWord",
	"agree":                    "Agree",
}

func GetSupportedNames() []string {
	supported := make([]string, 0)
	for key := range linkNameWithMatherName {
		supported = append(supported, key)
	}

	return supported
}

func IsCorrectNameOfMather(name string) bool {
	_, is := linkNameWithMatherName[name]

	return is
}

// ConvertNameToMatcher - .
func ConvertNameToMatcher(name string) string {
	return linkNameWithMatherName[name]
}
