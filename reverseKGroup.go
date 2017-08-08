package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	pivot := head
	var count int
	for count < k && pivot != nil {
		pivot = pivot.Next
		count++
	}
	if count == k {
		if pivot != nil {
			pivot = reverseKGroup(pivot, k)
		}
		for count > 0 {
			count--
			tmp := head.Next
			head.Next = pivot
			pivot = head
			head = tmp
		}
		head = pivot
	}
	return head
}
