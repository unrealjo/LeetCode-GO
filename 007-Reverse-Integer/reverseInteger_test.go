package reverseinteger

import "testing"

type TestCase struct {
	number  int
	reverse int
}

func TestReverse(t *testing.T) {
	ans := []TestCase{
		{123, 321},
		{-123, -321},
		{120, 21},
		{1534236469, 0},
		{0, 0},
		{-2147483648, 0},
	}
	for _, an := range ans {
		if an.reverse != reverse(an.number) {
			t.Error("Reverse not matching")
		}
	}
}
