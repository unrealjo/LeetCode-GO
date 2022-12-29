package romanToInteger

func romanToInt(s string) int {
	R := map[byte]int{
		'I': 1, 'X': 10, 'C': 100, 'M': 1000,
		'V': 5, 'L': 50, 'D': 500,
	}
	sum := 0
	length := len(s)
	var right byte
	var current byte
	for i := 0; i < length; i++ {
		current = s[i]
		if i+1 < length {
			right = s[i+1]
		}
		if R[right] > R[current] {
			sum -= R[current]
		} else {
			sum += R[current]
		}
	}
	return sum
}
