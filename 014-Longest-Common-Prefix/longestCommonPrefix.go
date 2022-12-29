package longestCommonPrefix

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var result strings.Builder
	Length := len(strs[0])
	strLen := 0
	// Find the smallest string length
	for _, str := range strs {
		strLen = len(str)
		if strLen < Length {
			Length = strLen
		}
	}
	state := true
	for i := 0; i < Length; i++ {
		current := strs[0][i]
		for _, str := range strs {
			if str[i] != current {
				state = false
				break
			}
		}
		// if all string have same character at same position
		if state {
			result.WriteByte(current)
		}
	}
	return result.String()
}
