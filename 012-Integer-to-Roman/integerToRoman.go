package integertoroman

import (
	"strings"
)

func intToRoman(num int) string {
	var result strings.Builder
	power := 1
	digit := 0
	numDigits := 0
	RomanCode := map[int][]string{
		1: {"I", "X", "C", "M"}, // 1 : 1, 10, 100, 1000
		4: {"IV", "XL", "CD"},   // 4 : 4, 40, 400
		5: {"V", "L", "D"},      // 5 : 5, 50, 500
		9: {"IX", "XC", "CM"},   // 9 : 9, 90, 900
	}
	// count number of digit in given number and its power
	for num/power != 0 {
		numDigits++
		power *= 10
	}
	// Extract digits and convert each digit to its  corresponding  power
	for i := numDigits; i >= 0; i-- {
		digit = num / power
		num -= digit * power
		power /= 10
		if digit >= 1 {
			if digit <= 3 {
				result.WriteString(strings.Repeat((RomanCode[1][i]), digit))
			} else if digit >= 6 && digit <= 8 {
				result.WriteString((RomanCode[5][i]))
				result.WriteString(strings.Repeat((RomanCode[1][i]), digit-5))
			} else {
				result.WriteString((RomanCode[digit][i]))
			}

		}
	}
	return result.String()
}
