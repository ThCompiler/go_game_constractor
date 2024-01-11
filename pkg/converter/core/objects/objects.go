package objects

import (
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/stringutilits"
)

type (
	Digit int8
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

type digitTriplet[T any] struct {
	Hundreds T
	Dozens   T
	Units    T
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
