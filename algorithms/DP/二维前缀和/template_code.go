package dp

/***
灵神的模版：
https://leetcode.cn/circle/discuss/UUuRex/
*/

/***
状态方程:
sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + a[i][j]
*/

type MatrixSum [][]int

func NewMatrixSum(matrix [][]int) MatrixSum {
	m := len(matrix)
	n := len(matrix[0])
	sum := make(MatrixSum, m+1)
	sum[0] = make([]int, n+1)
	// 注意一下顺序， i, j 都是正序
	for i:= range matrix{
		sum[i+1] = make([]int, n+1)
		for j := range matrix[0]{
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + matrix[i][j]
		}
	}
	return sum
}

/***
左上角  [r1,c1] 右下角 [r2,c2]  是闭区间。
 */
func (s MatrixSum) query(r1, c1, r2, c2 int) int {
	return s[r2+1][c2+1] - s[r2+1][c1] - s[r1][c2+1] + s[r1][c1]
}


