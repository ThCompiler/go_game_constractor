package words

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/declension"
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
