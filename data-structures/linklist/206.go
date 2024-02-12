package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/***
这也会写错， 真是服了你了啊
 */
func reverseList(head *ListNode) *ListNode {

	// 这是技巧，对吧， 可以写在一行
	var prev, cur *ListNode = nil, head

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	return prev
}
