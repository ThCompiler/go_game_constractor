package constants

type NumberType string

const (
	MaxIntegerPartLength = 306

	DecimalNumber    = NumberType(",")
	FractionalNumber = NumberType("/")

	CountDigits               = 10
	CountScaleNumberNameForms = 2
	CountNumberNameForms      = 3
)
