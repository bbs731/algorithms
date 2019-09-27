package C86

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

// 一点都不优雅， 能不能不要这多if 的特殊条件
func partition(head *ListNode, x int) *ListNode {
	var prev, cur, pos, tmp *ListNode
	prev, cur = head, head
	pos = nil

	for cur != nil {
		if cur.Val < x {
			// 第一个元素 head, head.Val < x
			if cur == head {
				pos = head
				prev = head
				cur = head.Next
				continue
			}

			// e.g [1,2,3,4,5]  x = 6
			if pos == prev {
				prev = cur
				pos = cur
				cur = cur.Next
				continue
			}

			prev.Next = cur.Next
			// prev does not change for this case as we need to move the cur element to pos->Next
			tmp = cur
			cur = cur.Next

			// insert tmp into head or proper position
			if pos == nil {
				tmp.Next = head
				head = tmp
				pos = head
			} else {
				tmp.Next = pos.Next
				pos.Next = tmp
				pos = tmp
			}

		} else {
			prev = cur
			cur = cur.Next
		}
	}
	return head
}
