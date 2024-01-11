package objects

import (
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/words"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/currency"
)

type CurrencyWords struct {
	Currencies map[words.CurrencyName]currency.Info `yaml:"currencies"`
}
