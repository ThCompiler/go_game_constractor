package matchers

var linkNameWithMatherName = map[string]string{
	"number":                "NumberMatcher",
	"positive_number":       "PositiveNumberMatcher",
	"positive_number_words": "PositiveNumberInWordsMatcher",
	"any":                   "AnyMatcher",
	"first_word":            "FirstWord",
	"agree":                 "Agree",
}

func GetSupportedNames() []string {
	supported := make([]string, 0)
	for key, _ := range linkNameWithMatherName {
		supported = append(supported, key)
	}

	return supported
}

func IsCorrectNameOfMather(name string) bool {
	_, is := linkNameWithMatherName[name]
	return is
}

func ConvertNameToMatcher(name string) string {
	return linkNameWithMatherName[name]
}
