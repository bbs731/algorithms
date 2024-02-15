package interview_questions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func rotateRight(head *ListNode, k int) *ListNode {
	n :=0
	for h:=head; h!=nil; h =h.Next{
		n++
	}
	if n == 0 || n==1 || k%n == 0 {
		return head
	}
	dummy := &ListNode{Next: head}
	k = k%n

	// rotate k times
	for i:=0; i<k; i++ {
		prev := dummy.Next // 怎么能这么难！ 这里不能再用 head 了，会更新
		for j:=1; j<=n-2; j++ {
			prev = prev.Next
		}
		tail := prev.Next
		prev.Next = nil
		tail.Next = dummy.Next
		dummy.Next = tail
	}
	return dummy.Next
}
