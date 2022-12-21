package zigzagconversion

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var res strings.Builder
	l := len(s)
	step := (numRows - 1) * 2
	for i := 0; i < numRows; i++ {
		// reset steps at last row
		for k := i; k < l; k += step { 
			res.WriteByte(s[k])
			if i < numRows-1 && i > 0 && k+step-i*2 < l {
				res.WriteByte(s[step+k-i*2])
			}

		}
	}
	return res.String()
}
