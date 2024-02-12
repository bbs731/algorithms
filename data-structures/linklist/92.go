package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/***
灵神的答案，太牛逼！，
重复默写100遍
*/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	p0 := dummy
	for i := 0; i < left-1; i++ {
		p0 = p0.Next
	}

	var pre, cur *ListNode = nil, p0.Next
	for i := 0; i < right-left+1; i++ {
		nxt := cur.Next
		cur.Next = pre // 每次循环只修改一个 Next，方便大家理解
		pre = cur
		cur = nxt
	}

	// 见视频
	p0.Next.Next = cur
	p0.Next = pre
	return dummy.Next
}

/***
链表的题目，怎么这么容易错啊， 好像，没办法一次写对啊
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	var pre, cur = dummy, head

	for step := 1; step < left; step++ {
		pre, cur = cur, cur.Next
	}
	saved_pre := pre

	pre, cur = dummy, head
	for step := 1; step < right; step++ {
		pre, cur = cur, cur.Next
	}
	// cur is now the end
	next := cur.Next
	cur.Next = nil
	l := reverseList(saved_pre.Next)
	saved_pre.Next.Next = next
	saved_pre.Next = l

	return dummy.Next
}

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
