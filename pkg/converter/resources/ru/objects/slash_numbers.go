package objects

import (
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/constants"
)

type SlashNumberUnitPrefixes struct {
	Units    [constants.CountDigits]string `yaml:"units"`
	Tens     [constants.CountDigits]string `yaml:"tens"`
	Dozens   [constants.CountDigits]string `yaml:"dozens"`
	Hundreds [constants.CountDigits]string `yaml:"hundreds"`
}
