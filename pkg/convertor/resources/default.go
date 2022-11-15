package resources

import (
    "github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/entity"
    "github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru"
    "github.com/ThCompiler/go_game_constractor/pkg/convertor/words/languages"
)

var supportResources = map[languages.Language]entity.Resources{
    languages.Russia: ru.GerResources(),
}

func IsKnowLanguages(lang languages.Language) bool {
    _, is := supportResources[lang]
    return is
}

func GetResources(lang languages.Language) entity.Resources {
    if IsKnowLanguages(lang) {
        return supportResources[lang]
    }
    return entity.Resources{}
}
