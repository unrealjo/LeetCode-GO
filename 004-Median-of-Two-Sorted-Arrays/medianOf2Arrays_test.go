package medianof2arrays

import (
	"testing"
)

type TestCase struct {
	arr1   []int
	arr2   []int
	result float64
}

func Test_median(t *testing.T) {
	ans := []TestCase{
		{
			[]int{1, 3},
			[]int{2},
			2.00,
		},
		{
			[]int{1, 2},
			[]int{4, 3},
			2.5,
		},
	}
	for _, val := range ans {
		med := median(val.arr1, val.arr2)
		if val.result != med {
			t.Errorf("Expected %v , Got %v", val.result, med)
		}
	}
}
