package euVatValidator

import (
	"reflect"
	"testing"

	"github.com/thereisnoplanb/validators/euVatValidator/enums/Status"
)

func TestValidator_Validate(t *testing.T) {
	type args struct {
		euVat                  string
		checkEuVatIsRegistered []bool
	}
	tests := []struct {
		name             string
		validator        *Validator
		args             args
		wantStatus       Status.Enum
		wantIsRegistered bool
		wantErr          bool
	}{
		{
			name:      "Valid number without checking registration in VIES.",
			validator: New(),
			args: args{
				euVat:                  "PL7251868136",
				checkEuVatIsRegistered: []bool{false},
			},
			wantStatus:       Status.Valid,
			wantIsRegistered: false,
			wantErr:          false,
		},
		{
			name:      "Valid number with checking registration in VIES",
			validator: New(),
			args: args{
				euVat:                  "PL7251868136",
				checkEuVatIsRegistered: []bool{true},
			},
			wantStatus:       Status.Valid,
			wantIsRegistered: true,
			wantErr:          false,
		},
		{
			name:      "Valid number without checking registration in VIES.",
			validator: New(),
			args: args{
				euVat:                  "PL0000000000",
				checkEuVatIsRegistered: []bool{false},
			},
			wantStatus:       Status.Valid,
			wantIsRegistered: false,
			wantErr:          false,
		},
		{
			name:      "Valid number with checking registration in VIES",
			validator: New(),
			args: args{
				euVat:                  "PL0000000000",
				checkEuVatIsRegistered: []bool{true},
			},
			wantStatus:       Status.Valid,
			wantIsRegistered: false,
			wantErr:          false,
		},
		{
			name:      "Invalid number - length",
			validator: New(),
			args: args{
				euVat:                  "PL",
				checkEuVatIsRegistered: nil,
			},
			wantStatus:       Status.IncorrectLength,
			wantIsRegistered: false,
			wantErr:          false,
		},
		{
			name:      "Invalid number - unrecognized country code",
			validator: New(),
			args: args{
				euVat:                  "XX0000000000",
				checkEuVatIsRegistered: nil,
			},
			wantStatus:       Status.UnrecognizedCountryCode,
			wantIsRegistered: false,
			wantErr:          false,
		},
		{
			name:      "Invalid number - format",
			validator: New(),
			args: args{
				euVat:                  "PL_123456789",
				checkEuVatIsRegistered: nil,
			},
			wantStatus:       Status.InvalidFormat,
			wantIsRegistered: false,
			wantErr:          false,
		},
		{
			name:      "Invalid number - checksum",
			validator: New(),
			args: args{
				euVat:                  "PL7251868137",
				checkEuVatIsRegistered: nil,
			},
			wantStatus:       Status.IncorrectChecksum,
			wantIsRegistered: false,
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStatus, gotIsRegistered, err := tt.validator.Validate(tt.args.euVat, tt.args.checkEuVatIsRegistered...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validator.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStatus, tt.wantStatus) {
				t.Errorf("Validator.Validate() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
			if gotIsRegistered != tt.wantIsRegistered {
				t.Errorf("Validator.Validate() gotIsRegistered = %v, want %v", gotIsRegistered, tt.wantIsRegistered)
			}
		})
	}
}
