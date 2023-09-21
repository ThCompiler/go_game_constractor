package ru

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/words"
	"github.com/pkg/errors"
)

var _ = words.Language(&Russian{})

type Russian struct {
	words *wordsConstantsForNumbers
}

func LoadRussianFromPath(path string) *Russian {
	wordsConstants, err := loadWordsConstantsFromFile(path)
	if err != nil {
		panic(errors.Wrapf(err, "could not load russian language"))
	}

	return &Russian{words: &wordsConstants.N2w}
}

func LoadRussian() *Russian {
	wordsConstants, err := loadWordsConstantsFromResources()
	if err != nil {
		panic(errors.Wrapf(err, "could not load russian language"))
	}

	return &Russian{words: &wordsConstants.N2w}
}

func (rs *Russian) GetMinusString() string {
	return rs.words.Sign.Minus
}
