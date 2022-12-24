package palindromenumber

func isPalindrome(x int) bool {
	k, reversedX := 1, 0
	sign := 1
	// check for negativity
	if x < 0 {
		sign = -1
		x = x * -1
	}
	if x < 10 {
		// if number is negative return false otherwise return true
		return !(sign == -1)
	}
	for x/k != 0 {
		// get reversedXd
		reversedX = reversedX*10 + x%(k*10)/k
		k *= 10
	}
	return sign*reversedX == x
}
