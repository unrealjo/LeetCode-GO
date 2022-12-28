package romanToInteger

import "testing"

type TestCase struct {
	romanVal string
	num      int
}

func TestRomanToInt(t *testing.T) {
	ans := []TestCase{
		{"III", 3},
		{"CIX", 109},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"MDCXCV", 1695},
	}
	for i, tc := range ans {
		if tc.num != romanToInt(tc.romanVal) {
			t.Errorf("Ivalid numbers at case %d", i+1)
		}
	}
}
