package ru

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/words"
)

// GetCurrencyAsWord Возвращает значение валюты в словах для указанной части числа
func (rs *Russian) GetCurrencyAsWord(numberInfo words.NumberInfo, numberAsTriplets []objects.RuneDigitTriplet) string {
	return ""
}

// GetCurrencyForFractionalNumber Возвращает значение валюты в случае если передана дробь
func (rs *Russian) GetCurrencyForFractionalNumber() string {
	return ""
}
