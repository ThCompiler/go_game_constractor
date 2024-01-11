package objects

import (
	"encoding/json"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/constants"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/objects"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/words"
	"github.com/ThCompiler/go_game_constractor/pkg/converter/resources/ru/objects/genders"
	"gopkg.in/yaml.v3"
)

type Sign struct {
	Minus string `yaml:"minus"`
}

type wordsByGender map[genders.Gender]string

type digitWord struct {
	word         string
	wordByGender wordsByGender
}

func (d *digitWord) WithGender() bool {
	return d.word == ""
}

func (d *digitWord) GetWordsByGender() wordsByGender {
	return d.wordByGender
}

func (d *digitWord) GetWord() string {
	return d.word
}

func (d *digitWord) UnmarshalYAML(n *yaml.Node) error {
	var err error
	if err = n.Decode(&d.word); err == nil {
		return nil
	}

	d.word = ""

	if err = n.Decode(&d.wordByGender); err == nil {
		return nil
	}

	d.wordByGender = nil

	return err
}

func (d *digitWord) UnmarshalJSON(b []byte) error {
	var err error
	if err = json.Unmarshal(b, &d.word); err == nil {
		return nil
	}

	d.word = ""

	if err = json.Unmarshal(b, &d.wordByGender); err == nil {
		return nil
	}

	d.wordByGender = nil

	return err
}

type DeclensionNumbers map[words.Declension][constants.CountDigits]digitWord

type DigitWords struct {
	Units    DeclensionNumbers `yaml:"units"`
	Tens     DeclensionNumbers `yaml:"tens"`
	Dozens   DeclensionNumbers `yaml:"dozens"`
	Hundreds DeclensionNumbers `yaml:"hundreds"`
}

type Words map[string]objects.Digit

type Digit struct {
	Units    Words
	Tens     Words
	Dozens   Words
	Hundreds Words
}

func (dn *DeclensionNumbers) convertToWords() Words {
	words := make(Words)

	for _, num2word := range *dn {
		for digit, word := range num2word {
			if word.WithGender() {
				genderWord := word.GetWordsByGender()

				for _, wrd := range genderWord {
					words[wrd] = objects.Digit(digit)
				}
			} else {
				words[word.GetWord()] = objects.Digit(digit)
			}
		}
	}

	return words
}

func NewWordsDigit(dw DigitWords) Digit {
	return Digit{
		Units:    dw.Units.convertToWords(),
		Tens:     dw.Tens.convertToWords(),
		Dozens:   dw.Dozens.convertToWords(),
		Hundreds: dw.Hundreds.convertToWords(),
	}
}