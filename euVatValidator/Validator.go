package euVatValidator

import "github.com/thereisnoplanb/vies"

// EU VAT validator.
type Validator struct {
	globalCheckEuVatIsRegisteredInVies bool
	viesClient                         *vies.Client
}
