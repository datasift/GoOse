package goose

import (
	"testing"
)

func TestIsValidLanguageCode(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected bool
	}{
		{"valid 2-letter code", "it", true},
		{"valid 3-letter code", "ita", true},
		{"reserved / invalid", "qaa", false},
		{"too long", "italia", false},
		{"missing", "", false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := isValidLanguageCode(tc.input)
			if actual != tc.expected {
				t.Errorf("Unexpected result from language validation: EXPECTED: %t, ACTUAL: %t", tc.expected, actual)
			}
		})
	}
}