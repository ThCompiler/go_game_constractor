package objects

import (
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/objects/declension"
)

const indexShift = 2

type declensionUnitName map[declension.Declension][constants.CountWordForms]string

type UnitScalesNames struct {
	Thousands      declensionUnitName `yaml:"thousands"`
	OtherEnding    declensionUnitName `yaml:"otherEnding"`
	OtherBeginning []string           `yaml:"otherBeginning"`
}

type WordScales map[string]uint64

type UnitScalesNamesToNumber struct {
	Words WordScales
}

func NewUnitScalesNamesToNumber(usn UnitScalesNames) UnitScalesNamesToNumber {
	usntn := UnitScalesNamesToNumber{
		Words: make(WordScales),
	}

	for _, num2word := range usn.Thousands {
		for _, word := range num2word {
			usntn.Words[word] = 1
		}
	}

	for i, wordBegin := range usn.OtherBeginning {
		for _, num2wordEnd := range usn.OtherEnding {
			for _, wordEnd := range num2wordEnd {
				usntn.Words[wordBegin+wordEnd] = uint64(i) + indexShift
			}
		}
	}

	return usntn
}
