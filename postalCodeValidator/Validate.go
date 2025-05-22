package postalCodeValidator

import (
	"regexp"

	"github.com/thereisnoplanb/globalization/CountryCode"
	"github.com/thereisnoplanb/validators/postalCodeValidator/enums/Status"
)

// Validates postal code against pattern for country.
//
// # Parameters:
//
//	countryCode CountryCode.Enum
//
// Country code, for example: PL.
//
//	postalCode string
//
// Postal code, for example: 00-001 for Warsaw, Poland.
//
// # Returns:
//
//	status Status.Enum
//
// Status of validation.
func (validator *Validator) Validate(countryCode CountryCode.Enum, postalCode string) (status Status.Enum) {

	if !countryCode.IsDefined() {
		return Status.UnrecognizedCountryCode
	}

	if pattern, ok := patterns[countryCode]; ok {
		if matched, err := regexp.MatchString(pattern, postalCode); err != nil || !matched {
			return Status.InvalidFormat
		}
	}

	return Status.Valid
}
