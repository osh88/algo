// Удалить n-й эл-т с конца
// Input:  [1,2,3,4,5], 2
// Output: [1,2,3,5]

package linked_list

func NewListNode[T comparable](items []T) *ListNode[T] {
	tail := &ListNode[T]{}
	head := tail

	for _, item := range items {
		tail.Next = &ListNode[T]{Value: item}
		tail = tail.Next
	}

	return head.Next
}

type ListNode[T comparable] struct {
	Value T
	Next  *ListNode[T]
}

func (o *ListNode[T]) Length() int {
	length := 0

	for node := o; node != nil; node = node.Next {
		length++
	}

	return length
}

func (o *ListNode[T]) ToSlice() []T {
	r := make([]T, 0, o.Length())

	for i := o; i != nil; i = i.Next {
		r = append(r, i.Value)
	}

	return r
}

// O(n)
// index: [1..N]
func (o *ListNode[T]) RemoveNthFromEnd(index int) *ListNode[T] {
	if index < 1 {
		return o
	}

	length := o.Length()

	if index > length {
		return o
	}

	dummy := &ListNode[T]{Next: o}
	node := dummy
	for i := 0; i < length-index; i++ {
		node = node.Next
	}

	if node.Next != nil {
		node.Next = node.Next.Next
	} else {
		node.Next = nil
	}

	return dummy.Next
}

func (o *ListNode[T]) Reverse() *ListNode[T] {
	var prev *ListNode[T]
	curr := o

	for curr != nil {
		tmp := curr
		curr = curr.Next
		tmp.Next = prev
		prev = tmp
	}

	return prev
}

// func (o *ListNode[T]) Middle() *ListNode[T] {
// 	r := o
// 	for range o.Length() / 2 {
// 		r = r.Next
// 	}
// 	return r
// }

func (o *ListNode[T]) Middle() *ListNode[T] {
	fast := o
	slow := o

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func (o *ListNode[T]) IsPalindrome() bool {
	left := o
	right := o.Middle().Reverse()

	for left != nil && right != nil {
		if left.Value != right.Value {
			return false
		}

		left = left.Next
		right = right.Next
	}

	return true
}
