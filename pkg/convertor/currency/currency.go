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

type Declensions map[declension.Declension][constants.CountScaleNumberNameForms]string

type NounGender struct {
	/**
	 * 0 => 'один', 1 => 'одна', 2 => 'одно'\
	 * Default: `0`
	 */
	Integer genders.Gender `yaml:"integer"`

	/**
	 * 0 => 'один', 1 => 'одна', 2 => 'одно'\
	 * Default: `1`
	 */
	FractionalPart genders.Gender `yaml:"fractionalPart"`
}

type CustomCurrency struct {
	/**
	 * Integer currency name forms\
	 * for digits [1, 2-4, 5-9]\
	 * e.g. ['рубль', 'рубля', 'рублей']
	 */
	CurrencyNameCases [constants.CountNumberNameForms]string `yaml:"currencyNameCases"`

	/**
	 * Integer currency name forms\
	 * for declensions
	 */
	CurrencyNameDeclensions Declensions `yaml:"currencyNameDeclensions"`

	/**
	 * Fractional number currency name forms\
	 * for digits [1, 2-4, 5-9]\
	 * e.g. ['копейка', 'копейки', 'копеек']
	 */
	FractionalPartNameCases [constants.CountNumberNameForms]string `yaml:"fractionalPartNameCases"`

	/**
	 * Fractional number currency name forms\
	 * for declensions
	 */
	FractionalPartNameDeclensions Declensions `yaml:"fractionalPartNameDeclensions"`

	CurrencyNounGender NounGender `yaml:"currencyNounGender"`
}
