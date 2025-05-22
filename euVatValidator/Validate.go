package euVatValidator

import (
	"regexp"

	"github.com/thereisnoplanb/validators/euVatValidator/enums/Status"
	"github.com/thereisnoplanb/vies/enums/ViesCountryCode"
)

// Validates EU VAT against length, pattern for EU country and checksum (if available).
//
// # Parameters:
//
//	euVat string
//
// EU VAT, for example: PL7251868136.
//
//	checkEuVatIsRegisteredInVies bool [OPTIONAL]
//
// If parameter checkEuVatIsRegisteredInVies is set or is ommited but global checkEuVatIsRegisteredInVies is set then
// after validation EU VAT is checked to be registered in VIES (VAT Information Exchange System)
//
// # Returns:
//
//	status Status.Enum
//
// Status of validation.
//
//	isRegisteredInVies bool
//
// If checkEuVatIsRegisteredInVies param or global checkEuVatIsRegisteredInVies is set, isRegisteredInVies stores information if EU VAT number is registered in VIES; false otherwise.
//
// # Errors:
//
//	err error
//
// Returns error when VIES vaildation fails.
//
// # Remarks
//
// If checkEuVatIsRegisteredInVies is ommited a global checkEuVatIsRegisteredInVies defined when creating new validator by New() method is used.
func (validator *Validator) Validate(euVat string, checkEuVatIsRegisteredInVies ...bool) (status Status.Enum, isRegisteredInVies bool, err error) {

	if len(euVat) < 3 {
		return Status.IncorrectLength, false, nil
	}

	viesCountryCode := ViesCountryCode.Enum(euVat[:2])

	if !viesCountryCode.IsDefined() {
		return Status.UnrecognizedCountryCode, false, nil
	}

	vatNumber := euVat[2:]

	if pattern, ok := patterns[viesCountryCode]; ok {
		if matched, err := regexp.MatchString(pattern, vatNumber); err != nil || !matched {
			return Status.InvalidFormat, false, err
		}
	}

	if checksum, ok := checksums[viesCountryCode]; ok {
		if checksum != nil && !checksum(vatNumber) {
			return Status.IncorrectChecksum, false, nil
		}
	}

	if (len(checkEuVatIsRegisteredInVies) > 0 && checkEuVatIsRegisteredInVies[0]) || (len(checkEuVatIsRegisteredInVies) == 0 && validator.globalCheckEuVatIsRegisteredInVies) {
		isRegisteredInVies, err = validator.viesClient.Validate(euVat)
	}
	return Status.Valid, isRegisteredInVies, err
}
