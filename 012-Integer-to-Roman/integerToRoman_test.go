package integertoroman

import "testing"

type TestCase struct {
	num      int
	romanVal string
}

func TestIntToRoman(t *testing.T) {
	ans := []TestCase{
		{3, "III"},
		{109, "CIX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}
	for i, tc := range ans {
		if tc.romanVal != intToRoman(tc.num) {
			t.Errorf("Ivalid roman value at case %d", i+1)
		}
	}
}
