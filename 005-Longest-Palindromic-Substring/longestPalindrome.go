package longestPalindrome

type Interval struct {
	vals []int
}

func (I Interval) len() int {
	return I.vals[1] - I.vals[0]
}
func GetMaxPalindrome(right, left int, s string) []int {
	l := len(s)
	for right < l && left >= 0 && s[right] == s[left] {
		left--
		right++
	}
	return []int{left + 1, right}
}
func longestPalindrome(s string) string {
	l1 := Interval{}
	l2 := Interval{}
	max := Interval{[]int{0, 0}}
	for i := range s {
		l1.vals = GetMaxPalindrome(i, i, s)
		l2.vals = GetMaxPalindrome(i+1, i, s)
		if l1.len() > l2.len() && l1.len() > max.len() {
			max = l1
		} else if l2.len() > max.len() {
			max = l2
		}
	}
	return s[max.vals[0]:max.vals[1]]
}
