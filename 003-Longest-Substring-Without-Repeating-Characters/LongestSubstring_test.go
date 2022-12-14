package LongestSubstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	type Testcase struct {
		str    string
		result int
	}
	ans := []Testcase{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}
	for _, an := range ans {
		if an.result != lengthOfLongestSubstring(an.str) {
			t.Error("Wrong length")
		}
	}
}
