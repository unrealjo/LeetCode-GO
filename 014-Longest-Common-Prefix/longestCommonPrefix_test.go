package longestCommonPrefix

import "testing"

type TestCase struct {
	strs    []string
	lPrefix string
}

func TestLongestCommonPrefix(t *testing.T) {
	tcs := []TestCase{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"butter", "buttermilk", "butterfly"}, "butter"},
	}
	for i, tc := range tcs {
		if tc.lPrefix != longestCommonPrefix(tc.strs) {
			t.Errorf("Incorrect prefix at test case %d", i+1)
		}
	}
}
