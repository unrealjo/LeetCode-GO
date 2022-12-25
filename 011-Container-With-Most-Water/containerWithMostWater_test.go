package containerwithmostwater

import "testing"

type Testcase struct {
	hight  []int
	expect int
}

func TestMaxArea(t *testing.T) {
	ans := []Testcase{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}
	for index, an := range ans {
		if an.expect != maxArea(an.hight) {
			t.Errorf("Wrong max area at case %d ", index+1)
		}
	}
}
