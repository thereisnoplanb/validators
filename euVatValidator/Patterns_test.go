package euVatValidator

import (
	"regexp"
	"testing"
)

func Test_pattens(t *testing.T) {
	for countryCode, pattern := range patterns {
		t.Run("Patterns", func(t *testing.T) {
			if _, err := regexp.Compile(pattern); err != nil {
				t.Errorf("Pattern %s for %s country code has errors.", pattern, countryCode)
			}
		})
	}
}
