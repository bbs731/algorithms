package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 /***
 还是会，错无数次， 面试过不去。
  */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	d1 := make([]int, 101)
	dl1 := 0
	for l := l1; l != nil; l = l.Next {
		//d1 = d1*10 + l.Val
		d1[dl1] = l.Val
		dl1++
	}
	d2 := make([]int, 101)
	dl2 := 0
	for l := l2; l != nil; l = l.Next {
		//d2 = d2*10 + l.Val
		d2[dl2] = l.Val
		dl2++
	}

	// reverse d1, d2
	for i, j := 0, dl1-1; i < j; i++ {
		d1[i], d1[j] = d1[j], d1[i]
		j--
	}
	for i, j := 0, dl2-1; i < j; i++ {
		d2[i], d2[j] = d2[j], d2[i]
		j--
	}

	dummy := &ListNode{}
	carry := 0

	for i := 0; i < 101; i++ {
		d := d1[i] + d2[i] + carry
		if d == 0 && i >= max(dl1, dl2) {   // 这里很难想。
			break
		}
		if d >= 10 {
			carry = 1
		} else {
			carry = 0   // 这里容易错。 忘记reset to 0
		}
		d %= 10
		node := &ListNode{d, dummy.Next}
		dummy.Next = node
	}
	return dummy.Next
}
