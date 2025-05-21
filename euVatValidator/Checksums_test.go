package euVatValidator

import (
	"testing"
)

func Test_checksumGermany(t *testing.T) {
	type args struct {
		vatNumber string
	}
	tests := []struct {
		name   string
		args   args
		wantOk bool
	}{
		{
			name: "Valid Mercedes-Benz AG EU VAT = DE321281763",
			args: args{
				vatNumber: "321281763",
			},
			wantOk: true,
		},
		{
			name: "Valid Siemens EU VAT = DE129274202",
			args: args{
				vatNumber: "129274202",
			},
			wantOk: true,
		},
		{
			name: "Invalid EU VAT = DE321281762",
			args: args{
				vatNumber: "321281762",
			},
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOk := checksumGermany(tt.args.vatNumber); gotOk != tt.wantOk {
				t.Errorf("checksumGermany() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_checksumPoland(t *testing.T) {
	type args struct {
		vatNumber string
	}
	tests := []struct {
		name   string
		args   args
		wantOk bool
	}{
		{
			name: "Valid Rainbow Tours EU VAT = PL7251868136",
			args: args{
				vatNumber: "7251868136",
			},
			wantOk: true,
		},
		{
			name: "Valid Zakład Automatyki i Urządzeń Precyzyjnych Time-Net EU VAT = PL7250013846",
			args: args{
				vatNumber: "7250013846",
			},
			wantOk: true,
		},
		{
			name: "Invalid EU VAT = PL7251868136",
			args: args{
				vatNumber: "7251868135",
			},
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOk := checksumPoland(tt.args.vatNumber); gotOk != tt.wantOk {
				t.Errorf("checksumPoland() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_checksumItaly(t *testing.T) {
	type args struct {
		vatNumber string
	}
	tests := []struct {
		name   string
		args   args
		wantOk bool
	}{
		{
			name: "Valid ORANGE S.R.L. EU VAT = IT03690530104",
			args: args{
				vatNumber: "03690530104",
			},
			wantOk: true,
		},
		{
			name: "Valid GUCCIO GUCCI S.P.A. EU VAT = IT04294710480",
			args: args{
				vatNumber: "04294710480",
			},
			wantOk: true,
		},
		{
			name: "Invalid EU VAT = IT03690530103",
			args: args{
				vatNumber: "03690530103",
			},
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOk := checksumItaly(tt.args.vatNumber); gotOk != tt.wantOk {
				t.Errorf("checksumItaly() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_checksumBelgium(t *testing.T) {
	type args struct {
		vatNumber string
	}
	tests := []struct {
		name   string
		args   args
		wantOk bool
	}{
		{
			name: "Valid SA ORANGE BELGIUM EU VAT = BE0456810810",
			args: args{
				vatNumber: "0456810810",
			},
			wantOk: true,
		},
		{
			name: "Valid SA GUCCI BELGIUM EU VAT = BE0449057639",
			args: args{
				vatNumber: "0449057639",
			},
			wantOk: true,
		},
		{
			name: "Invalid EU VAT = BE0456810811",
			args: args{
				vatNumber: "0456810811",
			},
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOk := checksumBelgium(tt.args.vatNumber); gotOk != tt.wantOk {
				t.Errorf("checksumBelgium() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
