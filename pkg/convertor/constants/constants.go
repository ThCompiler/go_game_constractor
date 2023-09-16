package constants

type NumberType string

const (
	DECIMAL_NUMBER    = NumberType(",")
	FRACTIONAL_NUMBER = NumberType("/")
)

const (
	MaxNumberPartLength = 306

	CountDigits          = 10
	CountWordForms       = 2
	CountNumberNameForms = 3

	UnitsScale    = 1
	ThousandScale = 2

	TripletSize = 3
)
