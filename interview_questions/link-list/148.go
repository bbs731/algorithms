package link_list

import "sort"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func sortList(head *ListNode) *ListNode {
	//dummy := &ListNode{}
	if head == nil {
		return nil
	}

	l := []*ListNode{}
	for h :=head; h!=nil; h=h.Next{
		l = append(l, h)
	}

	sort.Slice(l, func(i, j int) bool {
		return l[i].Val < l[j].Val
	})

	for i :=0; i<len(l)-1;i++ {
		l[i].Next = l[i+1]
	}
	l[len(l)-1].Next = nil
	return l[0]
}
