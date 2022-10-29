package errors

var linkNameWithErrorName = map[string]string{
	"positive_number": "PositiveNumberError",
	"number":          "NumberError",
}

func GetSupportedNames() []string {
	supported := make([]string, 0)
	for key, _ := range linkNameWithErrorName {
		supported = append(supported, key)
	}

	return supported
}

func IsCorrectNameOfError(name string) bool {
	_, is := linkNameWithErrorName[name]
	return is
}

func ConvertNameToError(name string) string {
	return linkNameWithErrorName[name]
}
