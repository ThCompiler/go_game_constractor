package ru

import (
	_ "embed" //nolint:golint //these are the rules for working with embed

	"github.com/ThCompiler/go_game_constractor/pkg/cleanenv"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/entity"
)

//go:embed currencies_strings.yml
var currenciesStrings []byte

//go:embed digit_words.yml
var digitWords []byte

//go:embed fractional_unit.yml
var fractionalUnit []byte

//go:embed ordinal_numbers.yml
var ordinalNumbers []byte

//go:embed sign.yml
var sign []byte

//go:embed slash_number_unit_prefixes.yml
var slashNumberUnitPrefixes []byte

//go:embed unit_scales_names.yml
var unitScalesNames []byte

func GerResources() entity.Resources {
	return entity.Resources{
		CurrenciesStrings:       currenciesStrings,
		DigitWords:              digitWords,
		FractionalUnit:          fractionalUnit,
		OrdinalNumbers:          ordinalNumbers,
		Sign:                    sign,
		SlashNumberUnitPrefixes: slashNumberUnitPrefixes,
		UnitScalesNames:         unitScalesNames,
		Ext:                     cleanenv.YAML,
	}
}
