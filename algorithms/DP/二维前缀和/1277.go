package dp

func countSquares(matrix [][]int) int {
	sum := NewMatrixSum(matrix)
	ans := 0
	for i:= range matrix {
		for j := range matrix[0] {
			if matrix[i][j] ==1 {
				for step :=1; sum.query(i,j, i+step-1, j+step-1) == step*step; step++{
					//ans = max(ans, step*step)
					ans +=1
				}
			}
		}
	}
	return ans
}


type MatrixSum [][]int

func NewMatrixSum(matrix [][]int) MatrixSum {
	m := len(matrix)
	n := len(matrix[0])
	sum := make([][]int, m+1)
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

func (s MatrixSum) query(r1, c1, r2, c2 int) int {
	m, n := len(s), len(s[0])
	if r2+1 >=m ||c2+1 >=n {
		return -1
	}
	return s[r2+1][c2+1] - s[r2+1][c1] - s[r1][c2+1] + s[r1][c1]
}


func max(a, b int ) int {
	if  a> b {
		return a
	}
	return b
}