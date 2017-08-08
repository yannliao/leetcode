package leetcode

import "container/heap"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKListsWithPQ(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}
	var pq priority

	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			pq = append(pq, lists[i])
		}
	}
	heap.Init(&pq)
	res := new(ListNode)
	t := res
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*ListNode)
		t.Next = item
		t = t.Next
		if item.Next != nil {
			heap.Push(&pq, item.Next)
		}
	}
	return res.Next
}

type priority []*ListNode

func (pq priority) Len() int { return len(pq) }

func (pq priority) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq priority) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priority) Push(x interface{}) {
	// n := len(*pq)
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

func (pq *priority) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
