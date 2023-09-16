package currency

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/declension"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/genders"
)

type Currency string

const (
	RUB    = Currency(`rub`)
	NUMBER = Currency(`number`)
	USD    = Currency(`usd`)
	EUR    = Currency(`eur`)
	CUSTOM = Currency(`custom`)
)

type Declensions map[declension.Declension][constants.CountWordForms]string

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

type CustomCurrency struct {
	/**
	 * IntegerPart currency name forms\
	 * for declensions
	 */
	IntegerPartNameDeclensions Declensions `yaml:"decimalCurrencyNameDeclensions"`

	/**
	 * Fractional number currency name forms\
	 * for declensions
	 */
	FractionalPartNameDeclensions Declensions `yaml:"fractionalPartNameDeclensions"`

	CurrencyNounGender NounGender `yaml:"currencyNounGender"`
}
