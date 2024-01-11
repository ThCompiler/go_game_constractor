package functions

import (
	"github.com/ThCompiler/go_game_constractor/pkg/converter/core/objects"
	"math"
	"strings"
)

const countPart = 3

func SplitNumberIntoThrees(number string) []objects.RuneDigitTriplet {
	// Сделать количество цифр числа кратным 3
	numberLength := len(number)
	numberScales := int(math.Ceil(float64(numberLength) / countPart))
	numberLengthGoal := numberScales * countPart
	lackOfDigits := numberLengthGoal - numberLength
	extendedNumber := strings.Repeat("0", lackOfDigits) + number

	r := []rune(extendedNumber)
	// Разделить число на классы по 3 цифры в каждом
	var cutNumber []objects.RuneDigitTriplet
	for i := 0; i < len(r); i += countPart {
		cutNumber = append(cutNumber, objects.RuneDigitTriplet{
			Hundreds: r[i],
			Dozens:   r[i+1],
			Units:    r[i+2],
		})
	}

	return cutNumber
}
