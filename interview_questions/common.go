package interview_questions


 type ListNode struct {
 	Val int
 	Next *ListNode
 }


 type Node struct {
 	Val int
 	Neighbors []*Node
 }


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <  b {
		return a
	}
	return b
}
