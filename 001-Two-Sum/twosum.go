package twosum

func twoSum(nums []int, target int) []int {
	l := len(nums)
	for index, n := range nums {
		for i := index; i < l; i++ {
			if nums[i]+n == target {
				return []int{index, i}
			}
		}
	}
	return []int{}
}
