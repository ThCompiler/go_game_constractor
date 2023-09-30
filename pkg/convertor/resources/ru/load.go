package ru

import (
	"bytes"
	"fmt"
	"github.com/ThCompiler/go_game_constractor/pkg/cleanenv"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/configs"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources/ru/objects"
	"github.com/pkg/errors"
	"os"
	"path"
	"path/filepath"
)

type wordConstants struct {
	N2w wordsConstantsForNumbers
	W2n wordsConstantsForWords
}

type wordsConstantsForNumbers struct {
	UnitScalesNames         objects.UnitScalesNames         `yaml:"unitScalesNames"`
	SlashNumberUnitPrefixes objects.SlashNumberUnitPrefixes `yaml:"slashNumberUnitPrefixes"`
	DigitWords              objects.DigitWords              `yaml:"digitWords"`
	CurrencyStrings         objects.CurrencyWords           `yaml:"currencyStrings"`
	FractionalUnit          objects.FractionalUnit          `yaml:"fractionalUnit"`
	OrdinalNumbers          objects.OrdinalNumbers          `yaml:"ordinalNumbers"`
	Sign                    objects.Sign                    `yaml:"sign"`
}

type wordsConstantsForWords struct {
	Digit                   objects.Digit
	UnitScalesNamesToNumber objects.UnitScalesNamesToNumber
}

func loadWordsConstantsFromFile(resourcesDirPath string) (*wordConstants, error) {
	if _, err := os.Stat(path.Join(resourcesDirPath)); err != nil {
		return nil, errors.Wrap(err, "error when try check config dir for russian resources")
	}

	wordConstants := wordConstants{
		N2w: wordsConstantsForNumbers{},
		W2n: wordsConstantsForWords{},
	}

	dir := filepath.Join(resourcesDirPath)

	// Sign
	err := cleanenv.ReadConfig(filepath.Join(dir, "sign.yml"), &wordConstants.N2w.Sign)
	if err != nil {
		return nil, fmt.Errorf("error load russian sign: %w", err)
	}

	// Digit words
	err = cleanenv.ReadConfig(filepath.Join(dir, "digit_words.yml"), &wordConstants.N2w.DigitWords)
	if err != nil {
		return nil, fmt.Errorf("error load russian digit words: %w", err)
	}

	// Fractional unit
	err = cleanenv.ReadConfig(filepath.Join(dir, "fractional_unit.yml"), &wordConstants.N2w.FractionalUnit)
	if err != nil {
		return nil, fmt.Errorf("error load russian fractional unit: %w", err)
	}

	// Unit scales names
	err = cleanenv.ReadConfig(filepath.Join(dir, "unit_scales_names.yml"), &wordConstants.N2w.UnitScalesNames)
	if err != nil {
		return nil, fmt.Errorf("error load russian unit scales names: %w", err)
	}

	// Slash number unit prefixes
	err = cleanenv.ReadConfig(
		filepath.Join(dir, "slash_number_unit_prefixes.yml"),
		&wordConstants.N2w.SlashNumberUnitPrefixes,
	)
	if err != nil {
		return nil, fmt.Errorf("error load russian slash number unit prefixes: %w", err)
	}

	// Currencies strings
	err = cleanenv.ReadConfig(filepath.Join(dir, "currency_strings.yml"), &wordConstants.N2w.CurrencyStrings)
	if err != nil {
		return nil, fmt.Errorf("error load russian currencies strings: %w", err)
	}

	// Ordinal numbers
	err = cleanenv.ReadConfig(filepath.Join(dir, "ordinal_numbers.yml"), &wordConstants.N2w.OrdinalNumbers)
	if err != nil {
		return nil, fmt.Errorf("error load russian ordinal numbers: %w", err)
	}

	wordConstants.W2n.Digit = objects.NewWordsDigit(wordConstants.N2w.DigitWords)
	wordConstants.W2n.UnitScalesNamesToNumber = objects.NewUnitScalesNamesToNumber(wordConstants.N2w.UnitScalesNames)

	return &wordConstants, nil
}

func loadWordsConstantsFromResources() (*wordConstants, error) {
	WordConstants := wordConstants{
		N2w: wordsConstantsForNumbers{},
		W2n: wordsConstantsForWords{},
	}

	res := configs.GetResources()

	// Sign
	err := cleanenv.ReadConfigFromReader(bytes.NewReader(res.Sign), res.Ext, &WordConstants.N2w.Sign)
	if err != nil {
		return nil, fmt.Errorf("error load russian sign: %w", err)
	}

	// Digit words
	err = cleanenv.ReadConfigFromReader(bytes.NewReader(res.DigitWords), res.Ext, &WordConstants.N2w.DigitWords)
	if err != nil {
		return nil, fmt.Errorf("error load russian digit words: %w", err)
	}

	// Fractional unit
	err = cleanenv.ReadConfigFromReader(bytes.NewReader(res.FractionalUnit), res.Ext, &WordConstants.N2w.FractionalUnit)
	if err != nil {
		return nil, fmt.Errorf("error load russian fractional unit: %w", err)
	}

	// Unit scales names
	err = cleanenv.ReadConfigFromReader(bytes.NewReader(res.UnitScalesNames), res.Ext,
		&WordConstants.N2w.UnitScalesNames)
	if err != nil {
		return nil, fmt.Errorf("error load russian unit scales names: %w", err)
	}

	// Slash number unit prefixes
	err = cleanenv.ReadConfigFromReader(bytes.NewReader(res.SlashNumberUnitPrefixes), res.Ext,
		&WordConstants.N2w.SlashNumberUnitPrefixes)
	if err != nil {
		return nil, fmt.Errorf("error load russian slash number unit prefixes: %w", err)
	}

	// Currencies strings
	err = cleanenv.ReadConfigFromReader(bytes.NewReader(res.CurrencyStrings), res.Ext,
		&WordConstants.N2w.CurrencyStrings)
	if err != nil {
		return nil, fmt.Errorf("error load russian currencies strings: %w", err)
	}

	// Ordinal numbers
	err = cleanenv.ReadConfigFromReader(bytes.NewReader(res.OrdinalNumbers), cleanenv.YAML,
		&WordConstants.N2w.OrdinalNumbers)
	if err != nil {
		return nil, fmt.Errorf("error load russian ordinal numbers: %w", err)
	}

	WordConstants.W2n.Digit = objects.NewWordsDigit(WordConstants.N2w.DigitWords)
	WordConstants.W2n.UnitScalesNamesToNumber = objects.NewUnitScalesNamesToNumber(WordConstants.N2w.UnitScalesNames)

	return &WordConstants, nil
}
