package twosum

/*
01-Given an array of integers `nums` and an integer `target`,
return indices of the two numbers such that they add up to `target.
*/
func twoSum(nums []int, target int) []int {
	var sum int
	dics := make(map[int]int)
	for i, v := range nums {
		sum = target - v
		_, ok := dics[sum]
		if ok {
			return []int{dics[sum], i}
		}
		dics[v] = i
	}
	return []int{}
}
