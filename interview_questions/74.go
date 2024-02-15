package interview_questions

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	var i int
	for i=0; i<m; i++ {
		if matrix[i][n-1] >= target {
			break
		}
	}
	if i == m {  // 这个判断，很重要啊！ 边界条件啊， 为啥总 panic 呢？
		return false
	}
	pos := sort.SearchInts(matrix[i], target)
	return matrix[i][pos]== target
}
