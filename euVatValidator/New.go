package euVatValidator

import "github.com/thereisnoplanb/vies"

// Creates new EU VAT validator.
//
// # Parameters:
//
//	checkEuVatIsRegisteredInVies bool [OPTIONAL]
//
// Sets global flag to check EU VAT is registered in VIES (VAT Information Exchange System) if checkEuVatIsRegisteredInVies is set to true.
//
// Default value is false.
//
// # Returns:
//
//	validator *Validator
//
// A new validator.
func New(checkEuVatIsRegisteredInVies ...bool) (validator *Validator) {
	return &Validator{
		globalCheckEuVatIsRegisteredInVies: len(checkEuVatIsRegisteredInVies) > 0 && checkEuVatIsRegisteredInVies[0],
		viesClient:                         vies.New(),
	}
}
