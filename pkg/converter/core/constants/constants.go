package constants

type NumberType string

const (
	DECIMAL_NUMBER    = NumberType(",")
	FRACTIONAL_NUMBER = NumberType("/")
)

type NumberPart int

const (
	INTEGER_PART    = NumberPart(0)
	FRACTIONAL_PART = NumberPart(1)
)

const (
	MaxNumberPartLength = 306

	TwoSignAfterRoundForCurrency = 2

	CountDigits = 10

	TripletSize = 3

	NoRoundIndicator = -1
)
