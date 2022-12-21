package palindromenumber

import "testing"

type TestCase struct {
	number int
	expect bool
}

func TestIsPalindrome(t *testing.T) {
	ans := []TestCase{
		{121, true},
		{-121, false},
		{-1, false},
		{0, true},
	}
	for i, an := range ans {
		if isPalindrome(an.number) != an.expect {
			t.Errorf("Case %d is not palindrome ", i+1)
		}
	}
}
