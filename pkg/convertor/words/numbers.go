package words

import (
	"fmt"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/currency"
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/words/languages"
	"github.com/ilyakaznacheev/cleanenv"
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
	WordsDigit              WordsDigit
	UnitScalesNamesToNumber UnitScalesNamesToNumber
}

func LoadWordsConstants(lang languages.Language, resourcesDirPath string) error {
	WordConstants = wordConstants{
		N2w: wordsConstantsForNumbers{},
		W2n: wordsConstantsForWords{},
	}

	// Sign
	err := cleanenv.ReadConfig(resourcesDirPath+"/"+string(lang)+"/sign.yml", &WordConstants.N2w.Sign)
	if err != nil {
		return fmt.Errorf("error load %s sign: %w", lang, err)
	}

	// Digit words
	err = cleanenv.ReadConfig(resourcesDirPath+"/"+string(lang)+"/digit_words.yml",
		&WordConstants.N2w.DigitWords)
	if err != nil {
		return fmt.Errorf("error load %s digit words: %w", lang, err)
	}

	// Fractional unit
	err = cleanenv.ReadConfig(resourcesDirPath+"/"+string(lang)+"/fractional_unit.yml",
		&WordConstants.N2w.FractionalUnit)
	if err != nil {
		return fmt.Errorf("error load %s fractional unit: %w", lang, err)
	}

	// Unit scales names
	err = cleanenv.ReadConfig(resourcesDirPath+"/"+string(lang)+"/unit_scales_names.yml",
		&WordConstants.N2w.UnitScalesNames)
	if err != nil {
		return fmt.Errorf("error load %s unit scales names: %w", lang, err)
	}

	// Slash number unit prefixes
	err = cleanenv.ReadConfig(resourcesDirPath+"/"+string(lang)+"/slash_number_unit_prefixes.yml",
		&WordConstants.N2w.SlashNumberUnitPrefixes)
	if err != nil {
		return fmt.Errorf("error load %s slash number unit prefixes: %w", lang, err)
	}

	// Currencies strings
	err = cleanenv.ReadConfig(resourcesDirPath+"/"+string(lang)+"/currencies_strings.yml",
		&WordConstants.N2w.CurrenciesStrings)
	if err != nil {
		return fmt.Errorf("error load %s currencies strings: %w", lang, err)
	}

	// Ordinal numbers
	err = cleanenv.ReadConfig(resourcesDirPath+"/"+string(lang)+"/ordinal_numbers.yml",
		&WordConstants.N2w.OrdinalNumbers)
	if err != nil {
		return fmt.Errorf("error load %s ordinal numbers: %w", lang, err)
	}

	WordConstants.W2n.WordsDigit = NewWordsDigit(WordConstants.N2w.DigitWords)
	WordConstants.W2n.UnitScalesNamesToNumber = NewUnitScalesNamesToNumber(WordConstants.N2w.UnitScalesNames)
	return nil
}
