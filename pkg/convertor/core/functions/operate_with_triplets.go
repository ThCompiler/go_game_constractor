package functions

import "github.com/ThCompiler/go_game_constractor/pkg/convertor/core/objects"

func RemoveZeroTripletFromBeginning(numberTripletArray []objects.RuneDigitTriplet) []objects.RuneDigitTriplet {
	for index, triplet := range numberTripletArray {
		if !triplet.IsZeros() {
			numberTripletArray = numberTripletArray[index:]
			break
		}
	}

	if len(numberTripletArray) == 1 && numberTripletArray[0].IsZeros() {
		return []objects.RuneDigitTriplet{}
	}

	return numberTripletArray
}

func IndexByEndOfLastNotZeroTriplet(numberTripletArray []objects.RuneDigitTriplet) int {
	res := -1

	for i := len(numberTripletArray) - 1; i >= 0; i-- {
		if !numberTripletArray[i].IsZeros() {
			res = len(numberTripletArray) - i - 1
			break
		}
	}

	return res
}
