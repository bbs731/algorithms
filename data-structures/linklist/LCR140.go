package linklist


/***
标准答案， 快慢指针
 */

func trainingPlan(head *ListNode, cnt int) *ListNode {
	fast, slow := head, head

	for i:=1; i<=cnt; i++{
		fast = fast.Next
	}

	for fast!= nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

