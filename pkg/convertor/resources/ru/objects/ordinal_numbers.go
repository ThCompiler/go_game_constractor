package objects

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/core/words"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects/genders"
)

type declensionOrdinalNumbersT map[words.Declension][constants.CountWordForms]string

type genderOrdinalNumbersT map[genders.Gender]declensionOrdinalNumbersT

type OrdinalNumbers struct {
	Units    []genderOrdinalNumbersT `yaml:"units,omitempty"`
	Tens     []genderOrdinalNumbersT `yaml:"tens,omitempty"`
	Dozens   []genderOrdinalNumbersT `yaml:"dozens,omitempty"`
	Hundreds []genderOrdinalNumbersT `yaml:"hundreds,omitempty"`
}
