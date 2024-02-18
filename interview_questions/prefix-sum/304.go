package prefix_sum

/***
二维前缀和

i+1, j+1
sum[i+1][j+1] =  sum[i+1][j] + sum[i][j+1] - sum[i][j] + nums[i][j]
 */

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])

	sums := make([][]int, m+1)
	for i := range sums {
		sums[i] = make([]int, n+1)
	}

	// build up 2D prefix sum
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sums[i+1][j+1] = sums[i+1][j] + sums[i][j+1] - sums[i][j] + matrix[i][j]
		}
	}

	return NumMatrix{
		sums,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// query 是一个闭区间 [(row1, col1),  (row2, col2)]
	//row1++
	//col1++
	//row2++
	//col2++

	return this.sums[row1][col1] - this.sums[row2+1][col1] - this.sums[row1][col2+1] + this.sums[row2+1][col2+1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
