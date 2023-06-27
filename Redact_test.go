package redact_test

import (
	"testing"

	"github.com/mvndaai/redact"
)

func TestWords(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{in: "", expected: ""},
		{in: "Bob", expected: "B**"},
		{in: "Bob Jones", expected: "B** J****"},
		{in: "Bob K Jones", expected: "B** K J****"},
		{in: "The 🐶🪵 is brown.", expected: "T** 🐶* i* b*****"},
		{in: "many   spaces", expected: "m***   s*****"},
		{in: "123 w 450 e", expected: "1** w 4** e"},
		{in: "220 Main Street", expected: "2** M*** S*****"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			actual := redact.Words(tt.in)
			if tt.expected != actual {
				t.Errorf("actual(%s) != expected(%s)", actual, tt.expected)
			}
		})
	}
}

func TestEmail(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{in: "", expected: ""},
		{in: "@", expected: "@"},
		{in: "example@example.com", expected: "e******@example.com"},
		{in: "🐶🪵@b.com", expected: "🐶*@b.com"},
		{in: "joe+s@m@gmail.com", expected: "j****@m@gmail.com"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			actual := redact.Email(tt.in)
			if tt.expected != actual {
				t.Errorf("actual(%s) != expected(%s)", actual, tt.expected)
			}
		})
	}
}

func TestPhone(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{in: "", expected: ""},
		{in: "1", expected: "*"},
		{in: "1234", expected: "****"},
		{in: "12345", expected: "1****"},
		{in: "1-2-34", expected: "1-2-**"},
		{in: "801-123-1234", expected: "801-123-****"},
		{in: "801.123.1234", expected: "801.123.****"},
		{in: "801.123.123", expected: "801.123****"},
		{in: "12🐶", expected: "***"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			actual := redact.Phone(tt.in)
			if tt.expected != actual {
				t.Errorf("actual(%s) != expected(%s)", actual, tt.expected)
			}
		})
	}
}

func TestWordOptions(t *testing.T) {
	tests := []struct {
		name        string
		in          string
		prefixChars uint
		suffixChars uint
		maxAsterisk uint
		expected    string
	}{
		{name: "Empty", in: "", prefixChars: 1, suffixChars: 0, expected: ""},
		{name: "One char", in: "a", prefixChars: 1, suffixChars: 1, expected: "a"},
		{name: "0 len", in: "abc", prefixChars: 0, suffixChars: 0, expected: "***"},
		{name: "First", in: "abc", prefixChars: 1, suffixChars: 0, expected: "a**"},
		{name: "Last", in: "abc", prefixChars: 0, suffixChars: 1, expected: "**c"},
		{name: "Long", in: "abcdefg", prefixChars: 4, suffixChars: 0, expected: "abcd***"},
		{name: "Long with max", in: "abcdefg", prefixChars: 4, suffixChars: 0, maxAsterisk: 2, expected: "abcd**"},
		{name: "Long reverse", in: "abcdefg", prefixChars: 0, suffixChars: 4, expected: "***defg"},
		{name: "Long reverse with max", in: "abcdefg", prefixChars: 0, suffixChars: 4, maxAsterisk: 2, expected: "**defg"},

		{name: "Too many short", in: "ab", prefixChars: 10, suffixChars: 10, expected: "a*"},
		{name: "Too many prefix", in: "abc", prefixChars: 10, suffixChars: 0, expected: "ab*"},
		{name: "Too many suffix", in: "abc", prefixChars: 0, suffixChars: 10, expected: "*bc"},
		{name: "Too many both", in: "abcdefg", prefixChars: 10, suffixChars: 10, expected: "abcde*g"},
		{name: "Too many after suffix", in: "abcdefg", prefixChars: 2, suffixChars: 10, expected: "ab*defg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := redact.WordOptions(tt.in, tt.prefixChars, tt.suffixChars, tt.maxAsterisk)
			if tt.expected != actual {
				t.Errorf("actual(%s) != expected(%s)", actual, tt.expected)
			}
		})
	}
}
