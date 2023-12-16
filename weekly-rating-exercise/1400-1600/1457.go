package _400_1600

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
 好题啊。 不需要 path, 只需要 mapCnt 记录 1~9 出现的次数
 
  */

 func isPalindrome(mapCnt map[int]int)bool{

 	odd := 0
 	for _, p := range mapCnt{
 		if p %2 !=0 {
 			odd++
 			if odd >1 {
 				return false
			}
		}
	}
 	return true
 }

func pseudoPalindromicPaths (root *TreeNode) int {
	// global path
	cnt := 0
	//path := make([]int, 0)
	mapCnt := make(map[int]int)

	var dfs func(*TreeNode)
	dfs = func(r *TreeNode){
		// leaf node
		if r.Left == nil && r.Right == nil {
			mapCnt[r.Val]++
			if isPalindrome(mapCnt) {
				cnt++
			}
			mapCnt[r.Val]--
			return
		}

		if r.Left != nil {
			//path = append(path, r.Val)
			mapCnt[r.Val]++
			dfs(r.Left)
			//path = path[:len(path)-1]
			mapCnt[r.Val]--
		}

		if r.Right !=nil {
			mapCnt[r.Val]++
			//path = append(path, r.Val)
			dfs(r.Right)
			//path = path[:len(path)-1]
			mapCnt[r.Val]--
		}
	}
	if root != nil {
		dfs(root)
	}
	return cnt
}
