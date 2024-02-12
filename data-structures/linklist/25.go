package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/***
太难了， 这面试， 肯定得挂啊！
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	p0 := dummy

	for true {
		p := p0
		for i := 0; i < k && p != nil; i++ {
			p = p.Next
		}

		if p == nil {
			return dummy.Next
		}

		var pre, cur *ListNode = nil, p0.Next

		for i := 0; i < k; i++ {
			nxt := cur.Next
			cur.Next = pre // 每次循环只修改一个 Next，方便大家理解
			pre = cur
			cur = nxt
		}

		// 见视频
		p0.Next.Next = cur
		tmp := p0.Next // 这个太重要了，想想，为什么？
		p0.Next = pre
		p0 = tmp
	}

	return dummy.Next
}

/***
https://leetcode.cn/problems/reverse-nodes-in-k-group/solutions/1992228/you-xie-cuo-liao-yi-ge-shi-pin-jiang-tou-plfs/
灵神的答案，真是简明啊！
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++ // 统计节点个数
	}

	dummy := &ListNode{Next: head}
	p0 := dummy
	var pre, cur *ListNode = nil, p0.Next
	for ; n >= k; n -= k {
		for i := 0; i < k; i++ {
			nxt := cur.Next
			cur.Next = pre // 每次循环只修改一个 Next，方便大家理解
			pre = cur
			cur = nxt
		}

		// 见视频
		nxt := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = nxt
	}
	return dummy.Next
}
