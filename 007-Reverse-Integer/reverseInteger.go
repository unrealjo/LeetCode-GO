package reverseinteger

import (
	"math"
)

func reverse(x int) int {
	result, k, sign := 0, 1, 1
	// reverse the number
	for x/k != 0 {
		/*
			Add the digits from last to first
			and multiply by 10 to avoid
			the addition between digits
		*/
		result = result*10 + x%(k*10)/k
		k *= 10
		// if 'result' is negative set sign to -1
		if result < 0 {
			sign = -1
		}
		// check if the number is not outside the signed 32-bit integer range
		if sign*result > math.MaxInt32 {
			return 0
		}
	}
	return result
}
