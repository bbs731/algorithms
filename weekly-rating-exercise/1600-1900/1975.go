package _600_1900

import "math"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 只有 1648 分数
// 这是一个贪心， 这就是纯粹的智商题啊。典型的面试题，能看出来就看出来，看不出来，觉得智商不够。
func maxMatrixSum(matrix [][]int) int64 {
	n := len(matrix)
	smallest := math.MaxInt / 2 // abs value
	sum := 0
	cnt := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum += abs(matrix[i][j])
			if smallest > abs(matrix[i][j]) {
				smallest = abs(matrix[i][j])
			}
			if matrix[i][j] < 0 {
				cnt++
			}
		}
	}
	if cnt%2 == 0 {
		return int64(sum)
	}
	return int64(sum - 2*smallest)
}
