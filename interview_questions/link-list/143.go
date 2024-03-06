package link_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	n := 0
	remain := 0
	var h *ListNode
	for h = head; h != nil; h = h.Next {
		n++
	}
	if n == 0 {
		return
	}

	if n&1 == 0 {
		remain = n/2 - 1
	} else {
		remain = n / 2
	}

	if remain == 0 {
		return
	}
	front := n - remain
	h = head
	for i := 1; i <= front-1; i++ {
		h = h.Next
	}

	rl := h.Next
	h.Next = nil

	//reverse(rl)
	//rdummy := ListNode{0, rl}
	var prev *ListNode
	prev, h = nil, rl

	for h != nil {
		next := h.Next
		h.Next = prev
		prev = h
		h = next
	}
	//rdummy.Next = prev
	//dummy := ListNode{0, head}

	h, rh := head, prev
	for i := 1; i <= remain; i++ {
		rn := rh.Next
		hn := h.Next

		h.Next = rh
		rh.Next = hn

		h = hn
		rh = rn
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 看到这种写法， 感觉自己弱爆了。
// 但是需要 额外 O（n)的空间的。
func reorderList(head *ListNode) {
	reorderListArr(head)
}

func reorderListArr(head *ListNode) {
	nodes := []*ListNode{}

	cur := head
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}

	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++

		if i == j {
			break
		}

		nodes[j].Next = nodes[i]
		j--
	}

	nodes[i].Next = nil
}
