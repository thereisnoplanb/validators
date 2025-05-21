package euVatValidator

import "github.com/thereisnoplanb/vies/enums/CountryCode"

var patterns = map[CountryCode.Enum]string{
	CountryCode.Austria:         `U[0-9]{8}`,
	CountryCode.Belgium:         `[0-1][0-9]{9}`,
	CountryCode.Bulgaria:        `[0-9]{9,10}`,
	CountryCode.Croatia:         `[0-9]{11}`,
	CountryCode.Cyprus:          `[0-9]{8}[A-Z]`,
	CountryCode.Czechia:         `[0-9]{8,10}`,
	CountryCode.Denmark:         `[0-9]{8}`,
	CountryCode.Estonia:         `[0-9]{9}`,
	CountryCode.Finland:         `[0-9]{8}`,
	CountryCode.France:          `[0-9A-Z]{2}[0-9]{9}`,
	CountryCode.Germany:         `[0-9]{9}`,
	CountryCode.Greece:          `[0-9]{9}`,
	CountryCode.Hungary:         `[0-9]{8}`,
	CountryCode.Ireland:         `[0-9]S[0-9]{5}[A-Z]|[0-9]{7}[A-W][A-I]`,
	CountryCode.Italy:           `[0-9]{11}`,
	CountryCode.Latvia:          `[0-9]{11}`,
	CountryCode.Lithuania:       `[0-9]{9}|[0-9]{12}`,
	CountryCode.Luxembourg:      `[0-9]{8}`,
	CountryCode.Malta:           `[0-9]{8}`,
	CountryCode.Netherlands:     `[0-9]{9}B[0-9]{2}`,
	CountryCode.NorthernIreland: `[0-9]{9}|[0-9]{12}|(GD|HA)[0-9]{3}`,
	CountryCode.Poland:          `[0-9]{10}`,
	CountryCode.Portugal:        `[0-9]{9}`,
	CountryCode.Romania:         `[0-9]{2,10}`,
	CountryCode.Slovakia:        `[0-9]{10}`,
	CountryCode.Slovenia:        `[0-9]{8}`,
	CountryCode.Spain:           `[A-Z][0-9]{7}[A-Z]|[0-9]{8}[A-Z]|[A-Z][0-9]{8}`,
	CountryCode.Sweden:          `[0-9]{12}`,
}
