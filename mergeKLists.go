package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	t := new(ListNode)
	res := t
	flag := true
	if len(lists) == 0 {
		return res.Next
	}
	for flag {
		j := 0
		for i := 0; i < len(lists); i++ {
			j = min(lists, j, i)
		}
		if lists[j] != nil {
			v := lists[j]
			lists[j] = v.Next
			v.Next = nil
			t.Next = v
			t = t.Next
		} else {
			flag = false
		}
	}
	return res.Next
}

func min(lists []*ListNode, j int, i int) int {
	if lists[j] == nil {
		return i
	}
	if lists[i] == nil {
		return j
	}
	if lists[j].Val > lists[i].Val {
		return i
	}
	return j
}
