package stringtointeger

import (
	"math"
	"strconv"
)

func myAtoi(s string) int {
	sign := 1
	sum := 0
	plusFound := false
	intParsed := 0
	for _, v := range s {
		if s == "" {
			return 0
		}
		//-------------- Parsing int -------------------
		if v >= '0' && v <= '9' {
			intParsed++
			n, err := strconv.ParseInt(string(v), 10, 32)
			if err == nil {
				sum = sum*10 + int(n)
			}
			// check for  -2^31 + 1 <= sum <= 2^31 -1
			if sum > math.MaxInt32 {
				if sign < 0 {
					return math.MinInt32
				} else {
					return math.MaxInt32
				}
			}
			// check if at last one int is parsed
		} else if sum >= 0 && intParsed > 0 {
			return sign * sum
		} else {
			//-------------- checking for edge case ---------------
			// check for ++ and +- case
			if plusFound || sign == -1 {
				return 0
			}
			if v == '-' {
				sign = -1
			} else if v == '+' {
				plusFound = true

			} else if v == ' ' {
				// to check for + serounded with spaces
				if plusFound {
					return 0
				}
				continue
			} else {
				break
			}
		}
	}
	return sign * sum
}
