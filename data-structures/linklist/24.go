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

/***
是不是有一百种写法， 每个人的思维都是不一样的。

但是犯错，也是有一百种样式！
 */

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		0,
		head,
	}

	// 灵神，引入的 p0 的这个概念，太牛了！
	var p0, prev, cur *ListNode = dummy, nil, head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next.Next
		next.Next = prev

		p0.Next.Next = cur
		p0.Next = next
		p0 = prev
	}

	return dummy.Next
}
