package entity

import "github.com/ThCompiler/go_game_constractor/pkg/cleanenv"

type Resources struct {
	CurrencyStrings         []byte
	DigitWords              []byte
	FractionalUnit          []byte
	OrdinalNumbers          []byte
	Sign                    []byte
	SlashNumberUnitPrefixes []byte
	UnitScalesNames         []byte
	Ext                     cleanenv.ConfigType
}
