package medianof2arrays

func median(a, b []int) float64 {
	var sum, i float64 = 0, 0
	for _, val := range append(a, b...) {
		sum += float64(val)
		i++
	}
	return sum / i
}
