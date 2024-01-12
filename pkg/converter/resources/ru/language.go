package ru

import (
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/words"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/objects/currency"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/objects/declension"
	"github.com/pkg/errors"
)

var _ = words.Language(&Russian{})

type Russian struct {
	words        *wordsConstantsForNumbers
	declension   declension.Declension
	currencyName words.CurrencyName
}

func LoadRussianFromPath(path string, opts ...Option) *Russian {
	wordsConstants, err := loadWordsConstantsFromFile(path)
	if err != nil {
		panic(errors.Wrapf(err, "could not load russian language"))
	}

	rs := &Russian{words: &wordsConstants.N2w}
	rs.SetOptions(opts...)
	return rs
}

func LoadRussian(opts ...Option) *Russian {
	wordsConstants, err := loadWordsConstantsFromResources()
	if err != nil {
		panic(errors.Wrapf(err, "could not load russian language"))
	}

	rs := &Russian{words: &wordsConstants.N2w}
	rs.SetOptions(opts...)
	return rs
}

func (rs *Russian) GetMinusString() string {
	if rs.words == nil {
		panic(ErrorLanguageNotLoaded)
	}
	return rs.words.Sign.Minus
}

func (rs *Russian) GetCurrency() currency.Info {
	if rs.words == nil {
		panic(ErrorLanguageNotLoaded)
	}

	return rs.words.CurrencyStrings.Currencies[rs.currencyName]
}

func (rs *Russian) SetOption(opt Option) {
	_ = opt(rs)
}

func (rs *Russian) SetOptions(opts ...Option) {
	if len(opts) == 0 {
		_ = DefaultOptions()(rs)
	}

	for _, opt := range opts {
		_ = opt(rs)
	}
}

// WithOptions Function create new instance of Russian with options. Words not copied.
func (rs *Russian) WithOptions(opts ...Option) *Russian {
	newRs := &Russian{
		words:        rs.words,
		declension:   rs.declension,
		currencyName: rs.currencyName,
	}

	newRs.SetOptions(opts...)

	return newRs
}

// IsCurrency Сообщает необходимо ли перевести число с указанием валюты
func (rs *Russian) IsCurrency() bool {
	return rs.currencyName != words.NUMBER
}

// IsNumber Сообщает необходимо ли перевести число как число без указание валюты
func (rs *Russian) IsNumber() bool {
	return rs.currencyName == words.NUMBER
}
