package twosum

import (
	"reflect"
	"testing"
)

type Testcase struct {
	nums   []int
	target int
	expect []int
}

func TestTwoSum(t *testing.T) {
	ans := []Testcase{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 5, -3}, 0, []int{0, 2}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for index, an := range ans {
		if !reflect.DeepEqual(twoSum(an.nums, an.target), an.expect) {
			t.Errorf("Wrong result at index %d", index)
		}
	}
}
