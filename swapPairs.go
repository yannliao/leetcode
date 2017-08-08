package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := head.Next.Next
	temp := head
	head = head.Next
	head.Next = temp
	temp.Next = swapPairs(last)
	return head
}
