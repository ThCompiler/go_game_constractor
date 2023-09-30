package constants

type NumberType string

const (
	DECIMAL_NUMBER    = NumberType(",")
	FRACTIONAL_NUMBER = NumberType("/")
)

const (
	MaxNumberPartLength = 306

	TwoSignAfterRoundForCurrency = 2

	CountDigits = 10

	TripletSize = 3
)
