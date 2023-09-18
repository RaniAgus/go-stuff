package morestrings

import (
	"testing"
)

func TestReverseRunes(test *testing.T) {
	cases := []struct {
		input, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, aCase := range cases {
		got := ReverseRunes(aCase.input)
		if got != aCase.want {
			test.Errorf("ReverseRunes(%q) == %q, want %q", aCase.input, got, aCase.want)
		}
	}
}
