package currency

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/words"
)

type Name string

const TwoSignAfterRound = 2

const (
	RUB    = Name(`rub`)
	NUMBER = Name(`number`)
	USD    = Name(`usd`)
	EUR    = Name(`eur`)
	CUSTOM = Name(`custom`)
)

type Declensions map[words.Declension][constants.CountWordForms]string

type NounGender struct {
	/**
	 * 0 => 'один', 1 => 'одна', 2 => 'одно'\
	 * Default: `0`
	 */
	IntegerPart words.Gender `yaml:"integerPart"`

	/**
	 * 0 => 'один', 1 => 'одна', 2 => 'одно'\
	 * Default: `1`
	 */
	FractionalPart words.Gender `yaml:"fractionalPart"`
}

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
