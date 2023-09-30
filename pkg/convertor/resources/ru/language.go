package ru

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/core/words"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/currency"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects/genders"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	femaleCurrency = words.CurrencyName(uuid.NewString())
)

var _ = words.Language(&Russian{})

type Russian struct {
	words *wordsConstantsForNumbers
}

func LoadRussianFromPath(path string) *Russian {
	wordsConstants, err := loadWordsConstantsFromFile(path)
	if err != nil {
		panic(errors.Wrapf(err, "could not load russian language"))
	}

	N2w := addFemaleCurrency(wordsConstants.N2w)
	return &Russian{words: &N2w}
}

func LoadRussian() *Russian {
	wordsConstants, err := loadWordsConstantsFromResources()
	if err != nil {
		panic(errors.Wrapf(err, "could not load russian language"))
	}
	N2w := addFemaleCurrency(wordsConstants.N2w)
	return &Russian{words: &N2w}
}

func addFemaleCurrency(wordsConstants wordsConstantsForNumbers) wordsConstantsForNumbers {
	wordsConstants.CurrencyStrings.Currencies[femaleCurrency] = currency.Info{
		CurrencyNounGender: currency.NounGender{
			IntegerPart:    genders.FEMALE,
			FractionalPart: genders.FEMALE,
		},
	}

	return wordsConstants
}

func (rs *Russian) GetMinusString() string {
	if rs.words == nil {
		panic(ErrorLanguageNotLoaded)
	}
	return rs.words.Sign.Minus
}

func (rs *Russian) GetCurrencyByName(name words.CurrencyName) currency.Info {
	if rs.words == nil {
		panic(ErrorLanguageNotLoaded)
	}

	return rs.words.CurrencyStrings.Currencies[name]
}
