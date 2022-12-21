package zigzagconversion

import "testing"

type Testcase struct {
	s       string
	numRows int
	expect  string
}

func TestConvert(t *testing.T) {
	ans := []Testcase{
		{"A", 1, "A"},
		{"timeisgood", 1, "timeisgood"},
		{"timeisgood", 2, "tmigoiesod"},
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"timeisgood", 5, "toiodmgesi"},
		{"timeisgood", 6, "tidmoeoigs"},
	}
	for i, an := range ans {
		res := convert(an.s, an.numRows)
		if an.expect != res {
			t.Errorf("Conversion fialed at index %d with result %s", i+1, res)
		}
	}
}
