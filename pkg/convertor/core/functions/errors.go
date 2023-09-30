package functions

import (
	"github.com/pkg/errors"
	"strings"
)

var (
	ErrorNumberContainsIncorrectChars = errors.New("converting number contains invalid characters for number")
	ErrorNumberHaveManySignChars      = errors.New("converting number contains not one sign character")
	ErrorNumberHaveSignCharNotInBegin = errors.New("converting number contains not beginning sign character")
	ErrorNumberHaveManyDelimiterChars = errors.New("converting number contains not one delimiter character")
)

func errorNumberContainsIncorrectChars(chars []string) error {
	return errors.Wrap(ErrorNumberContainsIncorrectChars, "with chars: \""+strings.Join(chars, ", ")+"\"")
}

func errorNumberHaveManyDelimiterChars(chars []string) error {
	return errors.Wrap(ErrorNumberHaveManyDelimiterChars, "with delimiter: \""+strings.Join(chars, ", ")+"\"")
}
