package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	pivot := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pivot.Next = l1
			pivot = pivot.Next
			l1 = l1.Next
		} else {
			pivot.Next = l2
			pivot = pivot.Next
			l2 = l2.Next
		}
	}
	for l1 != nil {
		pivot.Next = l1
		pivot = pivot.Next
		l1 = l1.Next
	}
	for l2 != nil {
		pivot.Next = l2
		pivot = pivot.Next
		l2 = l2.Next
	}
	return head.Next
}
