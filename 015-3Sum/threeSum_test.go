package threeSum

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums    []int
	triplet [][]int
}

func TestThreeSum(t *testing.T) {
	tcs := []TestCase{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			[]int{0, 1, 1}, [][]int{},
		},
		{
			[]int{3, 0, -2, -1, 1, 2},
			[][]int{{-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}},
		},
		{
			[]int{0, 0, 0}, [][]int{{0, 0, 0}},
		},
	}
	for i, tc := range tcs {

		if !reflect.DeepEqual(tc.triplet, threeSum(tc.nums)) {
			t.Errorf("Test failed at index %d", i+1)
		}
	}
}
