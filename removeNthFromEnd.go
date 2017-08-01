package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// var first *ListNode
	// var last *ListNode
	// t := ListNode{}
	first := new(ListNode)
	last := first
	first.Next, last.Next = head, head

	for i := 0; i < n && last.Next != nil; i++ {
		last = last.Next
	}
	for last.Next != nil {
		last = last.Next
		first = first.Next
	}

	var temp *ListNode
	temp = first.Next.Next
	first.Next.Next = nil
	if first.Next == head {
		head = temp
	}
	first.Next = temp
	// first.Next, first.Next.Next = first.Next.Next, nil

	return head
}
