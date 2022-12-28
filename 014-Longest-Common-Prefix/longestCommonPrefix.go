package longestCommonPrefix

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var result strings.Builder
	lenght := len(strs[0])
	strLen := 0
	// Find the smallest string lenght
	for _, str := range strs {
		strLen = len(str)
		if strLen < lenght {
			lenght = strLen
		}
	}
	state := true
	for i := 0; i < lenght; i++ {
		current := strs[0][i]
		for _, str := range strs {
			if str[i] != current {
				state = false
				break
			}
		}
		// if all string have same charachater at same position
		if state {
			result.WriteByte(current)
		}
	}
	return result.String()
}
