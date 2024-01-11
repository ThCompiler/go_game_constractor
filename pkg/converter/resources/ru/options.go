package ru

import (
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/words"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/objects/declension"
)

type Option func(*Russian) *Russian

func DefaultOptions() Option {
	return func(russian *Russian) *Russian {
		russian.currencyName = words.NUMBER
		russian.declension = declension.NOMINATIVE
		return russian
	}
}

func WithCurrency(name words.CurrencyName) Option {
	return func(russian *Russian) *Russian {
		russian.currencyName = name
		return russian
	}
}

func AsNumber() Option {
	return func(russian *Russian) *Russian {
		russian.currencyName = words.NUMBER
		return russian
	}
}

func WithDeclension(declension words.Declension) Option {
	return func(russian *Russian) *Russian {
		russian.declension = declension
		return russian
	}
}
