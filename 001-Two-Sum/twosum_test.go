package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}

	if !reflect.DeepEqual(twoSum(nums, 9), []int{0, 1}) {
		t.Error("Failed, two sum")
	}

	if !reflect.DeepEqual(twoSum(nums, 6), []int{}) {
		t.Error("Failed, two sum")
	}
}
