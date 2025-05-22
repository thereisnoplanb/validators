package euVatValidator

import "github.com/thereisnoplanb/vies/enums/ViesCountryCode"

type checksum func(euVat string) (ok bool)

var checksums = map[ViesCountryCode.Enum]checksum{
	ViesCountryCode.Austria:         nil,
	ViesCountryCode.Belgium:         nil,
	ViesCountryCode.Bulgaria:        nil,
	ViesCountryCode.Croatia:         nil,
	ViesCountryCode.Cyprus:          nil,
	ViesCountryCode.Czechia:         nil,
	ViesCountryCode.Denmark:         nil,
	ViesCountryCode.Estonia:         nil,
	ViesCountryCode.Finland:         nil,
	ViesCountryCode.France:          nil,
	ViesCountryCode.Germany:         checksumGermany,
	ViesCountryCode.Greece:          nil,
	ViesCountryCode.Hungary:         nil,
	ViesCountryCode.Ireland:         nil,
	ViesCountryCode.Italy:           checksumItaly,
	ViesCountryCode.Latvia:          nil,
	ViesCountryCode.Lithuania:       nil,
	ViesCountryCode.Luxembourg:      nil,
	ViesCountryCode.Malta:           nil,
	ViesCountryCode.Netherlands:     nil,
	ViesCountryCode.NorthernIreland: nil,
	ViesCountryCode.Poland:          checksumPoland,
	ViesCountryCode.Portugal:        nil,
	ViesCountryCode.Romania:         nil,
	ViesCountryCode.Slovakia:        nil,
	ViesCountryCode.Slovenia:        nil,
	ViesCountryCode.Spain:           nil,
	ViesCountryCode.Sweden:          nil,
}

func checksumPoland(vatNumber string) (ok bool) {
	if len(vatNumber) != 10 {
		return false
	}
	runes := []rune(vatNumber)
	weights := []int{6, 5, 7, 2, 3, 4, 5, 6, 7}
	sum := 0
	for i := range 9 {
		sum += int(runes[i]-'0') * weights[i]
	}
	return (sum % 11) == int(runes[9]-'0')
}

func checksumGermany(vatNumber string) (ok bool) {
	if len(vatNumber) != 9 {
		return false
	}
	runes := []rune(vatNumber)
	sum := 0
	for i := range 8 {
		value := (sum + int(runes[i]-'0')) % 10
		if value == 0 {
			value = 10
		}
		sum = (2 * value) % 11
	}
	sum = 11 - sum
	return sum == int(runes[8]-'0')
}

func checksumItaly(vatNumber string) (ok bool) {
	if len(vatNumber) != 11 {
		return false
	}
	runes := []rune(vatNumber)
	sum := 0
	for i, r := range runes {
		if i%2 == 0 {
			sum += int(r - '0')
		} else {
			value := 2 * int(r-'0')
			if value > 9 {
				value = value - 9
			}
			sum += value
		}
	}
	return sum%10 == 0
}

func checksumBelgium(vatNumber string) (ok bool) {
	if len(vatNumber) != 10 {
		return false
	}
	runes := []rune(vatNumber)
	sum := 0
	for i := range 8 {
		sum = 10*sum + int(runes[i]-'0')
	}
	return (97 - (sum % 97)) == int(10*(runes[8]-'0')+(runes[9]-'0'))
}

// func checksumItaly2(vat string) bool {
// 	if len(vat) != 11 {
// 		return false
// 	}
// 	y, err := strconv.Atoi(vat[7:10])
// 	if err != nil {
// 		return false
// 	}
// 	if y < 1 || (y > 100 && y != 120 && y != 121 && y != 999) {
// 		return false
// 	}

// 	if vat[:7] == "0000000" {
// 		return false
// 	}
// 	return checkDigit(vat)
// }

// func checkDigit(vat string) bool {
// 	sum := sumOfFigures(vat[:10])
// 	chk := 10 - (sum % 10)
// 	if chk == 10 {
// 		chk = 0
// 	}
// 	return chk == int(vat[10]-'0')
// }

// func sumOfFigures(vat string) int {
// 	sum := 0
// 	reverse := true
// 	for i := len(vat) - 1; i >= 0; i-- {
// 		num := int(vat[i] - '0')
// 		if reverse {
// 			num *= 2
// 			if num > 9 {
// 				num -= 9
// 			}
// 		}
// 		sum += num
// 		reverse = !reverse
// 	}
// 	return sum
// }
