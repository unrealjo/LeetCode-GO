package LongestSubstring

func lengthOfLongestSubstring(s string) int {
	charList := map[rune]int{}
	left := 0
	max := 0
	for right, v := range s {
		n, ok := charList[v]
		if ok && n > left {
			left = n
		}
		if max < right-left+1 {
			max = right - left + 1
		}
		charList[v] = right + 1
	}
	return max
}
