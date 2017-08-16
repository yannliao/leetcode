package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKListsRecursion(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	} else if length == 1 {
		return lists[0]
	}
	l1 := mergeKListsRecursion(lists[:length/2])
	l2 := mergeKListsRecursion(lists[length/2:])
	res := mergeTwoLists(l1, l2)
	return res
}
