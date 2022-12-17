package medianof2arrays

import "sort"

func findMedianSortedArrays(a, b []int) float64 {
	var list []float64
	for _, v := range a {
		list = append(list, float64(v))
	}
	for _, v := range b {
		list = append(list, float64(v))
	}
	sort.Float64s(list)
	l := len(list)

	/*
		Note : number elements is n but in the implementation is equal to l - 1
		If the total number elements given is odd, then the median is:
		median = [(n/2)th element + (n/2)th element] / 2
	*/
	if l%2 == 0 {
		return (list[l/2] + list[l/2-1]) / 2.0
	}
	/*
		If the total number elements given is odd, then the median is:
		median = (n+1/2)th element
	*/
	return float64(list[l/2])
}
