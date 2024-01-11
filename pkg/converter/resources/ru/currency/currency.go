package currency

import (
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/words"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/objects/genders"
)

const (
	RUB = words.CurrencyName(`rub`)
	USD = words.CurrencyName(`usd`)
	EUR = words.CurrencyName(`eur`)
)

type Declensions map[words.Declension][constants.CountWordForms]string

type Info struct {
	/**
	 * IntegerPart currency name forms\
	 * for declensions
	 */
	DecimalCurrencyNameDeclensions Declensions `yaml:"decimalCurrencyNameDeclensions"`

	/**
	 * Fractional number currency name forms\
	 * for declensions
	 */
	FractionalPartNameDeclensions Declensions `yaml:"fractionalPartNameDeclensions"`

	CurrencyNounGender NounGender `yaml:"currencyNounGender"`
}

type NounGender struct {
	/**
	 * 0 => 'один', 1 => 'одна', 2 => 'одно'\
	 * Default: `0`
	 */
	IntegerPart genders.Gender `yaml:"integerPart"`

	/**
	 * 0 => 'один', 1 => 'одна', 2 => 'одно'\
	 * Default: `1`
	 */
	FractionalPart genders.Gender `yaml:"fractionalPart"`
}
