package LongestSubstring

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	max, strLen := 0, 0
	str, char := "", ""
	l := len([]rune(s))
	for i := 0; i < l; i++ {
		str = string(s[i])
		for _, c := range s[i:] {
			char = string(c)
			if !strings.Contains(str, char) {
				str += char
			} else {
				str = char
			}
			strLen = len([]rune(str))
			if strLen > max {
				max = strLen
			}
		}
	}
	return max
}
