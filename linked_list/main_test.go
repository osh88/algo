// go test *.go

package linked_list

import (
	"slices"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	list := NewListNode([]int{1, 2, 3, 4, 5})

	list = list.RemoveNthFromEnd(2)

	assertSlicesEqual(t, []int{1, 2, 3, 5}, list.ToSlice(), "")
}

func TestReverse(t *testing.T) {
	list := NewListNode([]int{1, 2, 3, 4, 5})

	list = list.Reverse()

	assertSlicesEqual(t, []int{5, 4, 3, 2, 1}, list.ToSlice(), "")
}

func TestMiddle(t *testing.T) {
	type Test struct {
		Input  []int
		Output []int
	}
	tests := []Test{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}},
		{[]int{1, 2}, []int{2}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		list := NewListNode(test.Input)

		list = list.Middle()

		assertSlicesEqual(t, test.Output, list.ToSlice(), "")
	}
}

func TestIsPalindrome(t *testing.T) {
	type Test struct {
		Input  []int
		Output bool
	}
	tests := []Test{
		{[]int{1, 2, 3, 4, 5}, false},
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2, 3, 2, 1}, true},
		{[]int{}, true},
		{[]int{1, 1}, true},
		{[]int{1, 2, 1}, true},
	}

	for i, test := range tests {
		list := NewListNode(test.Input)

		if act := list.IsPalindrome(); test.Output != act {
			t.Errorf("[%d] exp != act, %v != %v", i, test.Output, act)
		}
	}
}

func assertSlicesEqual[S ~[]E, E comparable](t *testing.T, expected, actual S, message string) {
	if !slices.Equal(expected, actual) {
		if message != "" {
			t.Errorf("%s: exp: %v, act: %v", message, expected, actual)
		} else {
			t.Errorf("exp != act, %v != %v", expected, actual)
		}
	}
}
