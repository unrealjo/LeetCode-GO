package threeSum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	length := len(nums)
	var H, C, L int
	/*

		C : current value
		H : high value
		L : low value
		I preferd to store repeated operations in variables
	*/
	for i := 0; i < length-2; i++ {
		C = nums[i]
		if i == 0 || (i > 0 && C != nums[i-1]) {
			low := i + 1
			high := length - 1
			for low < high {
				L = nums[low]
				H = nums[high]
				if L+H+C == 0 {
					result = append(result, []int{C, L, H})
					for low < high && nums[low] == nums[low+1] {
						low++
					}
					for low < high && nums[high] == nums[high-1] {
						high--
					}
					low++
					high--
				} else if L+H+C > 0 {
					high--
				} else {
					low++
				}
			}
		}
	}
	return result
}
