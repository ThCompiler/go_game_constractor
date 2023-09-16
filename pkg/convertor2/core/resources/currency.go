package resources

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor2/core/currency"
)

var supportCurrency = map[currency.Name]currency.Info{}

func AddCurrency(currency currency.Name, customCurrency currency.Info) {
	supportCurrency[currency] = customCurrency
}

func IsKnowCurrency(curc currency.Name) bool {
	_, is := supportCurrency[curc]

	return is
}

func GetCurrency(curc currency.Name) currency.Info {
	if IsKnowCurrency(curc) {
		return supportCurrency[curc]
	}

	return currency.Info{}
}
