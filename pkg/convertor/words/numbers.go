package words

import (
	"bytes"
	"fmt"
	"github.com/ThCompiler/go_game_constractor/pkg/cleanenv"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/currency"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/resources"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/languages"
	"github.com/pkg/errors"
	"os"
	"path"
	"path/filepath"
)

type wordConstants struct {
	N2w wordsConstantsForNumbers
	W2n wordsConstantsForWords
}

var WordConstants wordConstants

type CurrencyWords struct {
	Currencies map[currency.Currency]currency.CustomCurrency `yaml:"currencies"`
}

type Sign struct {
	Minus string `yaml:"minus"`
}

type wordsConstantsForNumbers struct {
	UnitScalesNames         UnitScalesNames         `yaml:"unitScalesNames"`
	SlashNumberUnitPrefixes SlashNumberUnitPrefixes `yaml:"slashNumberUnitPrefixes"`
	DigitWords              DigitWords              `yaml:"digitWords"`
	CurrenciesStrings       CurrencyWords           `yaml:"currenciesStrings"`
	FractionalUnit          FractionalUnit          `yaml:"fractionalUnit"`
	OrdinalNumbers          OrdinalNumbers          `yaml:"ordinalNumbers"`
	Sign                    Sign                    `yaml:"sign"`
}

type wordsConstantsForWords struct {
	Digit                   Digit
	UnitScalesNamesToNumber UnitScalesNamesToNumber
}

func LoadWordsConstants(lang languages.Language, resourcesDirPath string) error {
	if _, err := os.Stat(path.Join(resourcesDirPath, string(lang))); err != nil {
		if !os.IsNotExist(err) {
			return errors.Wrap(err, "error when try check config dir for convertor lib")
		}

		return LoadWordsConstantsFromResources(lang)
	}

	return LoadWordsConstantsFromFile(lang, resourcesDirPath)
}

func LoadWordsConstantsFromFile(lang languages.Language, resourcesDirPath string) error {
	WordConstants = wordConstants{
		N2w: wordsConstantsForNumbers{},
		W2n: wordsConstantsForWords{},
	}

	dir := filepath.Join(resourcesDirPath, string(lang))

	// Sign
	err := cleanenv.ReadConfig(filepath.Join(dir, "sign.yml"), &WordConstants.N2w.Sign)
	if err != nil {
		return fmt.Errorf("error load %s sign: %w", lang, err)
	}

	// Digit words
	err = cleanenv.ReadConfig(filepath.Join(dir, "digit_words.yml"), &WordConstants.N2w.DigitWords)
	if err != nil {
		return fmt.Errorf("error load %s digit words: %w", lang, err)
	}

	// Fractional unit
	err = cleanenv.ReadConfig(filepath.Join(dir, "fractional_unit.yml"), &WordConstants.N2w.FractionalUnit)
	if err != nil {
		return fmt.Errorf("error load %s fractional unit: %w", lang, err)
	}

	// Unit scales names
	err = cleanenv.ReadConfig(filepath.Join(dir, "unit_scales_names.yml"), &WordConstants.N2w.UnitScalesNames)
	if err != nil {
		return fmt.Errorf("error load %s unit scales names: %w", lang, err)
	}

	// Slash number unit prefixes
	err = cleanenv.ReadConfig(filepath.Join(dir, "slash_number_unit_prefixes.yml"), &WordConstants.N2w.SlashNumberUnitPrefixes)
	if err != nil {
		return fmt.Errorf("error load %s slash number unit prefixes: %w", lang, err)
	}

	// Currencies strings
	err = cleanenv.ReadConfig(filepath.Join(dir, "currencies_strings.yml"), &WordConstants.N2w.CurrenciesStrings)
	if err != nil {
		return fmt.Errorf("error load %s currencies strings: %w", lang, err)
	}

	// Ordinal numbers
	err = cleanenv.ReadConfig(filepath.Join(dir, "ordinal_numbers.yml"), &WordConstants.N2w.OrdinalNumbers)
	if err != nil {
		return fmt.Errorf("error load %s ordinal numbers: %w", lang, err)
	}

	WordConstants.W2n.Digit = NewWordsDigit(WordConstants.N2w.DigitWords)
	WordConstants.W2n.UnitScalesNamesToNumber = NewUnitScalesNamesToNumber(WordConstants.N2w.UnitScalesNames)

	return nil
}

func LoadWordsConstantsFromResources(lang languages.Language) error {
	if !resources.IsKnowLanguages(lang) {
		return ErrorUnknownLanguage
	}

	WordConstants = wordConstants{
		N2w: wordsConstantsForNumbers{},
		W2n: wordsConstantsForWords{},
	}

	res := resources.GetResources(lang)

	// Sign
	err := cleanenv.ReadConfigFromReader(bytes.NewReader(res.Sign), res.Ext, &WordConstants.N2w.Sign)
	if err != nil {
		return fmt.Errorf("error load %s sign: %w", lang, err)
	}

	// Digit words
	err = cleanenv.ReadConfigFromReader(bytes.NewReader(res.DigitWords), res.Ext, &WordConstants.N2w.DigitWords)
	if err != nil {
		return fmt.Errorf("error load %s digit words: %w", lang, err)
	}

	// Fractional unit
	err = cleanenv.ReadConfigFromReader(bytes.NewReader(res.FractionalUnit), res.Ext, &WordConstants.N2w.FractionalUnit)
	if err != nil {
		return fmt.Errorf("error load %s fractional unit: %w", lang, err)
	}

	// Unit scales names
	err = cleanenv.ReadConfigFromReader(bytes.NewReader(res.UnitScalesNames), res.Ext,
		&WordConstants.N2w.UnitScalesNames)
	if err != nil {
		return fmt.Errorf("error load %s unit scales names: %w", lang, err)
	}

	// Slash number unit prefixes
	err = cleanenv.ReadConfigFromReader(bytes.NewReader(res.SlashNumberUnitPrefixes), res.Ext,
		&WordConstants.N2w.SlashNumberUnitPrefixes)
	if err != nil {
		return fmt.Errorf("error load %s slash number unit prefixes: %w", lang, err)
	}

	// Currencies strings
	err = cleanenv.ReadConfigFromReader(bytes.NewReader(res.CurrenciesStrings), res.Ext,
		&WordConstants.N2w.CurrenciesStrings)
	if err != nil {
		return fmt.Errorf("error load %s currencies strings: %w", lang, err)
	}

	// Ordinal numbers
	err = cleanenv.ReadConfigFromReader(bytes.NewReader(res.OrdinalNumbers), cleanenv.YAML,
		&WordConstants.N2w.OrdinalNumbers)
	if err != nil {
		return fmt.Errorf("error load %s ordinal numbers: %w", lang, err)
	}

	WordConstants.W2n.Digit = NewWordsDigit(WordConstants.N2w.DigitWords)
	WordConstants.W2n.UnitScalesNamesToNumber = NewUnitScalesNamesToNumber(WordConstants.N2w.UnitScalesNames)

	return nil
}
