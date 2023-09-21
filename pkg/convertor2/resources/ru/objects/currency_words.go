package objects

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/words"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/resources/ru/currency"
)

type CurrencyWords struct {
	Currencies map[words.CurrencyName]currency.Info `yaml:"currencies"`
}
