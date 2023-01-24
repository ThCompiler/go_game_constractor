package cleanenv

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

var (
	ErrorFieldNotProvided       = errors.New("field is required but the value is not provided")
	ErrorWrongTypeOfField       = errors.New("wrong type of field")
	ErrorInvalidMapItem         = errors.New("invalid map item")
	ErrorNotSupportedFileFormat = errors.New("this extation format doesn't supported by the parser")
	ErrorNotSupportedType       = errors.New("this type of field doesn't supported by the parser")
)

func errorFieldNotProvided(fieldName string) error {
	return errors.Wrap(ErrorFieldNotProvided, "with field`s name: \""+fieldName+"\"")
}

func errorWrongTypeOfField(kind reflect.Kind) error {
	return errors.Wrap(ErrorWrongTypeOfField, "with type: \""+kind.String()+"\"")
}

func errorInvalidMapItem(item string) error {
	return errors.Wrap(ErrorInvalidMapItem, "with item: \""+item+"\"")
}

func errorNotSupportedFileFormatAsString(format string) error {
	return errors.Wrap(ErrorNotSupportedFileFormat, fmt.Sprintf("with extation format: '%s'", format))
}

func errorNotSupportedFileFormatAsConfigType(format ConfigType) error {
	return errors.Wrap(ErrorNotSupportedFileFormat, fmt.Sprintf("with extation format: '%d'", format))
}

func errorNotSupportedType(pkg, tp string) error {
	return errors.Wrap(ErrorNotSupportedType, fmt.Sprintf("with type: '%s.%s'", pkg, tp))
}
