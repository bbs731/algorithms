package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/***
第一遍，还是写错， 而且还不知道错在哪
 */
func swapPairs(head *ListNode) *ListNode {

	dummy := &ListNode{Next: head}
	prev, cur := dummy, head

	for cur != nil {
		next := cur.Next
		if next == nil {
			break
		}

		nn := next.Next
		next.Next = cur
		cur.Next = nn
		prev.Next = next

		//reset // now cur has taken the place of next
		// 这种题目， 太他妈容易错了。
		prev = cur
		cur = prev.Next
	}

	return dummy.Next
}
