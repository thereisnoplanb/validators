package euVatValidator

import (
	"github.com/thereisnoplanb/vies/enums/ViesCountryCode"
)

var patterns = map[ViesCountryCode.Enum]string{
	ViesCountryCode.Austria:         `U[0-9]{8}`,
	ViesCountryCode.Belgium:         `[0-1][0-9]{9}`,
	ViesCountryCode.Bulgaria:        `[0-9]{9,10}`,
	ViesCountryCode.Croatia:         `[0-9]{11}`,
	ViesCountryCode.Cyprus:          `[0-9]{8}[A-Z]`,
	ViesCountryCode.Czechia:         `[0-9]{8,10}`,
	ViesCountryCode.Denmark:         `[0-9]{8}`,
	ViesCountryCode.Estonia:         `[0-9]{9}`,
	ViesCountryCode.Finland:         `[0-9]{8}`,
	ViesCountryCode.France:          `[0-9A-Z]{2}[0-9]{9}`,
	ViesCountryCode.Germany:         `[0-9]{9}`,
	ViesCountryCode.Greece:          `[0-9]{9}`,
	ViesCountryCode.Hungary:         `[0-9]{8}`,
	ViesCountryCode.Ireland:         `[0-9]S[0-9]{5}[A-Z]|[0-9]{7}[A-W][A-I]`,
	ViesCountryCode.Italy:           `[0-9]{11}`,
	ViesCountryCode.Latvia:          `[0-9]{11}`,
	ViesCountryCode.Lithuania:       `[0-9]{9}|[0-9]{12}`,
	ViesCountryCode.Luxembourg:      `[0-9]{8}`,
	ViesCountryCode.Malta:           `[0-9]{8}`,
	ViesCountryCode.Netherlands:     `[0-9]{9}B[0-9]{2}`,
	ViesCountryCode.NorthernIreland: `[0-9]{9}|[0-9]{12}|(GD|HA)[0-9]{3}`,
	ViesCountryCode.Poland:          `[0-9]{10}`,
	ViesCountryCode.Portugal:        `[0-9]{9}`,
	ViesCountryCode.Romania:         `[0-9]{2,10}`,
	ViesCountryCode.Slovakia:        `[0-9]{10}`,
	ViesCountryCode.Slovenia:        `[0-9]{8}`,
	ViesCountryCode.Spain:           `[A-Z][0-9]{7}[A-Z]|[0-9]{8}[A-Z]|[A-Z][0-9]{8}`,
	ViesCountryCode.Sweden:          `[0-9]{12}`,
}
