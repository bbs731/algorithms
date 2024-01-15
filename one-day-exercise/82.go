package one_day_exercise

/*

给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。



示例 1：


输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]
示例 2：


输入：head = [1,1,1,2,3]
输出：[2,3]

 */

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

// 哎！ 太容易错了， 链表的问题，怎么办呢？

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	prev := dummy
	cur := head

	for cur != nil {
		next := cur.Next
		for next != nil && cur.Val == next.Val {
			next = next.Next
			cur = cur.Next
		}
		if prev.Next == cur {
			// move forward
			prev = cur
			cur = next

		} else {
			// remove cur node
			prev.Next = next
			cur = next
		}
	}
	return dummy.Next
}
