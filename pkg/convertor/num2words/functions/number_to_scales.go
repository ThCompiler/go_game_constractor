package functions

import (
	"github.com/ThCompiler/go_game_constractor/pkg/convertor/num2words/objects"
	"math"
	"strings"
)

func NumberToScales(number string) []objects.RuneDigitTriplet {
	// Сделать количество цифр числа кратным 3
	numberLength := len(number)
	numberScales := int(math.Ceil(float64(numberLength) / 3.0))
	numberLengthGoal := numberScales * 3
	lackOfDigits := numberLengthGoal - numberLength
	extendedNumber := strings.Repeat("0", lackOfDigits) + number

	r := []rune(extendedNumber)
	// Разделить число на классы по 3 цифры в каждом
	var cutNumber []objects.RuneDigitTriplet
	for i := 0; i < len(r); i += 3 {
		cutNumber = append(cutNumber, objects.RuneDigitTriplet{
			Hundreds: r[i],
			Dozens:   r[i+1],
			Units:    r[i+2],
		})
	}
	return cutNumber
}
