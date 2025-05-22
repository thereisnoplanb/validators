package PostalCodeValidator

import (
	"reflect"
	"testing"

	"github.com/thereisnoplanb/globalization/CountryCode"
	"github.com/thereisnoplanb/validators/postalCodeValidator/enums/Status"
)

func TestValidator_Validate(t *testing.T) {
	type args struct {
		countryCode CountryCode.Enum
		postalCode  string
	}
	tests := []struct {
		name       string
		validator  *Validator
		args       args
		wantStatus Status.Enum
	}{
		{
			name:      "Valid postal code.",
			validator: New(),
			args: args{
				countryCode: CountryCode.Poland,
				postalCode:  "98-100",
			},
			wantStatus: Status.Valid,
		},
		{
			name:      "Invalid postal code.",
			validator: New(),
			args: args{
				countryCode: CountryCode.Poland,
				postalCode:  "AA-AAA",
			},
			wantStatus: Status.InvalidFormat,
		},
		{
			name:      "Invalid postal code - unrecognized country code",
			validator: New(),
			args: args{
				countryCode: "XX",
				postalCode:  "98-100",
			},
			wantStatus: Status.UnrecognizedCountryCode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStatus := tt.validator.Validate(tt.args.countryCode, tt.args.postalCode)
			if !reflect.DeepEqual(gotStatus, tt.wantStatus) {
				t.Errorf("Validator.Validate() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
