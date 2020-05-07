package discogs

import "testing"

func TestCurrency(t *testing.T) {
	tests := []struct {
		currency string
		expected string
		err      error
	}{
		{currency: "", expected: "USD", err: nil},
		{currency: "USD", expected: "USD", err: nil},
		{currency: "GBP", expected: "GBP", err: nil},
		{currency: "EUR", expected: "EUR", err: nil},
		{currency: "CAD", expected: "CAD", err: nil},
		{currency: "AUD", expected: "AUD", err: nil},
		{currency: "JPY", expected: "JPY", err: nil},
		{currency: "CHF", expected: "CHF", err: nil},
		{currency: "MXN", expected: "MXN", err: nil},
		{currency: "BRL", expected: "BRL", err: nil},
		{currency: "NZD", expected: "NZD", err: nil},
		{currency: "SEK", expected: "SEK", err: nil},
		{currency: "ZAR", expected: "ZAR", err: nil},
		{currency: "RUR", expected: "", err: ErrCurrencyNotSupported},
	}

	for _, testCase := range tests {
		cur, err := currency(testCase.currency)
		if err != testCase.err {
			t.Errorf("Currency error = %v, expected %v", err, testCase.err)
		}
		if cur != testCase.expected {
			t.Errorf("Currency = %v, expected %v", cur, testCase.expected)
		}
	}
}
