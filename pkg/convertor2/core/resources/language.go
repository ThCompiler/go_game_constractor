package resources

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/words"
)

var supportResources = map[words.LanguageName]words.Language{}

func IsKnowLanguages(lang words.LanguageName) bool {
	_, is := supportResources[lang]

	return is
}

func GetResources(lang words.LanguageName) words.Language {
	if IsKnowLanguages(lang) {
		return supportResources[lang]
	}

	return nil
}
