package stringtointeger

import "testing"

type TestCase struct {
	s      string
	expect int
}

func TestMyAtoi(t *testing.T) {
	ans := []TestCase{
		{"", 0},
		{" ", 0},
		{"42", 42},
		{"ABS", 0},
		{"9-42", 9},
		{"3.14", 3},
		{"-42", -42},
		{"a3.14", 0},
		{"  -89  ", -89},
		{"21474836460", 2147483647},
	}
	for i, an := range ans {
		if myAtoi(an.s) != an.expect {
			t.Errorf("Atoi failed at test case %d", i+1)
		}
	}
}
