package interview_questions

/****
这题目，现场是想不出来的啊！咋能锻炼一下呢？
 */
func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		// 这里是 i 和 n-1-i 互换，容易搞错啊！
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
