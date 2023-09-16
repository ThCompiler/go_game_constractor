package objects

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/stringutilits"
)

type (
	WordForm   int8
	Digit      int8
	NumberForm int8 // change word ending by digit
)

const (
	SINGULAR_WORD = WordForm(0)
	PLURAL_WORD   = WordForm(1)
)

const (
	FIRST_FORM  = NumberForm(0) // when digit is one (1)
	SECOND_FORM = NumberForm(1) // when digit between 2 and 4
	THIRD_FORM  = NumberForm(2) // when digit between 4 and 9
)

type Number struct {
	Sign       string
	FirstPart  string
	Divider    constants.NumberType
	SecondPart string
}

type ResultNumberT struct {
	Sign           string
	FirstPart      string
	FirstPartName  string
	SecondPart     string
	SecondPartName string
}

type ConvertedScalesToWords struct {
	Result          string
	UnitNameForm    NumberForm
	LastScaleIsZero bool
}

type digitTriplet[T any] struct {
	Units    T
	Dozens   T
	Hundreds T
}

type (
	RuneDigitTriplet    digitTriplet[rune]
	StringDigitTriplet  digitTriplet[string]
	NumericDigitTriplet digitTriplet[Digit]
)

func (ndt NumericDigitTriplet) IsZeros() bool {
	return ndt.Units == 0 && ndt.Dozens == 0 && ndt.Hundreds == 0
}

func (ndt NumericDigitTriplet) ToRune() RuneDigitTriplet {
	return RuneDigitTriplet{
		Units:    stringutilits.ToRune(int8(ndt.Units)),
		Dozens:   stringutilits.ToRune(int8(ndt.Dozens)),
		Hundreds: stringutilits.ToRune(int8(ndt.Hundreds)),
	}
}

func (rdt RuneDigitTriplet) IsZeros() bool {
	return rdt.Units == '0' && rdt.Dozens == '0' && rdt.Hundreds == '0'
}

func (rdt RuneDigitTriplet) ToNumeric() NumericDigitTriplet {
	return NumericDigitTriplet{
		Units:    Digit(stringutilits.ToDigit(rdt.Units)),
		Dozens:   Digit(stringutilits.ToDigit(rdt.Dozens)),
		Hundreds: Digit(stringutilits.ToDigit(rdt.Hundreds)),
	}
}
