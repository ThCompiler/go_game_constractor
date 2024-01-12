package objects

import (
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/objects/declension"
)

type DeclensionFractionalUnits map[declension.Declension][constants.CountWordForms]string

type fractionalUnitsDeclensionsT struct {
	Tens     DeclensionFractionalUnits `yaml:"tens"`
	Hundreds DeclensionFractionalUnits `yaml:"hundreds"`
}

type FractionalUnit struct {
	FractionalUnitsDeclensions fractionalUnitsDeclensionsT            `yaml:"fractionalUnitsDeclensions"`
	FractionalUnitsBeginning   []string                               `yaml:"fractionalUnitsBeginning"`
	FractionalUnitPrefixes     [constants.CountNumberNameForms]string `yaml:"fractionalUnitPrefixes"`
	FractionalUnitEndings      DeclensionFractionalUnits              `yaml:"fractionalUnitEndings"`
}
