package addtwonumbers

import (
	"reflect"
	"testing"
)

type testCase struct {
	arr1     []int
	arr2     []int
	expected []int
}

func (Node *ListNode) add(data []int) {
	for _, value := range data {
		temp := &ListNode{value, nil}
		if Node == nil {
			Node = temp
		} else {
			iter := Node
			for iter.Next != nil {
				iter = iter.Next
			}
			iter.Next = temp
		}
	}
}
func (Node *ListNode) array() []int {
	list := []int{}
	iter := Node.Next
	for iter != nil {
		list = append(list, iter.Val)
		iter = iter.Next
	}
	return list
}

func Test_addTwoNumbers(t *testing.T) {
	res := []testCase{
		{
			[]int{2, 4, 3},
			[]int{5, 6, 4},
			[]int{7, 0, 8},
		},
		{
			[]int{0},
			[]int{0},
			[]int{0},
		},
		{
			[]int{1, 2, 3},
			[]int{3, 6, 8},
			[]int{4, 8, 1, 1},
		},
	}
	for _, answer := range res {
		l1 := &ListNode{}
		l1.add(answer.arr1)
		l2 := &ListNode{}
		l2.add(answer.arr2)
		l3 := addTwoNumbers(l1, l2)
		if !reflect.DeepEqual(answer.expected, l3.array()) {
			t.Errorf("Expected %v but got %v", answer.expected, l3.array())
		}
	}
}
