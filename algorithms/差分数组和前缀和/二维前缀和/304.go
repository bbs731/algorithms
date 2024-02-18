package dp

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
	// query 是一个闭区间 [(row1, col1),  (row2, col2)],  还有 [ ) 左闭右开的区间， 看灵神的模板。
	// 这里之前考虑错了， 以 （0，0）左上角，作为基准的前缀和
	//row1++
	//col1++
	//row2++
	//col2++

	return this.sums[row1][col1] - this.sums[row2+1][col1] - this.sums[row1][col2+1] + this.sums[row2+1][col2+1]
}
