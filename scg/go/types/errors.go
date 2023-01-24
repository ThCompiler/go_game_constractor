package types

import "github.com/pkg/errors"

var ErrorNotSupportedType = errors.New("not supported type " +
	"fro converting from string. please add type with function AddType")
