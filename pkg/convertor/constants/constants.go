package constants

type NumberType string

const (
	MaxIntegerPartLength = 306

	DecimalNumber    = NumberType(",")
	FractionalNumber = NumberType("/")

	CountDigits               = 10
	CountScaleNumberNameForms = 2
	CountNumberNameForms      = 3

	UnitsScale    = 1
	ThousandScale = 2

	TripletSize = 3
)
