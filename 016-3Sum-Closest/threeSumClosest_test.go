package threeSumClosest

import (
	"testing"
)

type TestCase struct {
	nums    []int
	target  int
	closest int
}

func TestThreeSum(t *testing.T) {
	tcs := []TestCase{
		{[]int{1, 5, 4, 8, 9}, 11, 10},
		{[]int{-1, 2, 1, -4}, 1, 2},
	}
	for i, tc := range tcs {

		if tc.closest != threeSumClosest(tc.nums, tc.target) {
			t.Errorf("Wrong closest at index %d", i+1)
		}
	}
}
