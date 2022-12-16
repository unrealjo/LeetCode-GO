package longestPalindrome

import "testing"

type Testcase struct {
	str    string
	result string
}

func TestLongestPalindrome(t *testing.T) {
	ans := []Testcase{
		{"cbbd", "bb"},
		{"babad", "aba"},
		{"cac", "cac"},
		{"aa", "aa"},
		{"", ""},
		{"dallam", ""},
		{"scabsbacz", "cabsbac"},
	}

	for index, an := range ans {
		if an.result != longestPalindrome(an.result) {
			t.Errorf("Failed at testcase %d", index+1)
		}
	}
}
