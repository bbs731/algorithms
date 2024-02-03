package dp

import "fmt"

/***

f[i][j] = max(sum[j+1]-sum[i+1] - f[i+1][j],  sum[j]-sum[i] - f[i][j-1])
 */


func stoneGameVII(stones []int) int {
	n := len(stones)
	sum := make([]int, n+1)
	for i:=1;i<n+1; i++ {
		sum[i]= sum[i-1] + stones[i-1]
	}

	//f := make([][][2]int, n)
	f := make([][]int, n)
	for i :=range f {
		f[i] = make([]int, n)
	}

	for i:=n-2; i>=0; i-- {
		for j:=i+1; j<n; j++ {
			f[i][j] = max(sum[j+1]-sum[i+1] - f[i+1][j], sum[j]-sum[i]-f[i][j-1])
		}
	}
	return f[0][n-1]
}


/***
这个计算，就是错误的， 为什么呢？ 搞不懂，为什么是错误的。

f[i][j][1] = max(f[i+1][j][0] + sum[j+1]-sum[i+1], f[i][j-1][0] + sum[j] - sum[i])
f[i][j][0] = min(f[i+1][j][1], f[i][j-1][1])

f[i+1][j][0] = min( f[i+1][j-1][1]， f[i+2][j][1])
f[i][j-1][0] = min( f[i+1][j-1][1], f[i][j-2][1])
 */
func stoneGameVII(stones []int) int {
	n := len(stones)
	sum := make([]int, n+1)
	for i:=1;i<n+1; i++ {
		sum[i]= sum[i-1] + stones[i-1]
	}

	f := make([][][2]int, n)
	for i :=range f {
		f[i] = make([][2]int, n)
	}

	for i:=n-2; i>=0; i-- {
		for j:=i+1; j<n; j++ {
			a , b := sum[j+1]-sum[i+1] - f[i+1][j][1], sum[j]-sum[i] - f[i][j-1][1]
			if a > b {
				// 那么 [i][j][1] 会选择 ith element 因此
				f[i][j][0] = f[i+1][j][1]
			} else if a < b {
				// 会选择 jth element  因此，bob 只能选择
				f[i][j][0] = f[i][j-1][1]
			} else {
				 // a == b
				 f[i][j][0] = min(f[i+1][j][1], f[i][j-1][1])
			}

			f[i][j][1] = max(f[i+1][j][0] + sum[j+1]-sum[i+1], f[i][j-1][0] + sum[j]-sum[i])

			//if i+2 <n {
			//	f[i+1][j][0] = min(f[i+1][j-1][1], f[i+2][j][1])
			//}
			//if j-2 >=0 {
			//	f[i][j-1][0] = min(f[i+1][j-1][1], f[i][j-2][1])
			//}
			//f[i][j][0] = min(f[i+1][j][1], f[i][j-1][1])
		}
	}
	fmt.Println(f)
	return f[0][n-1][1] - f[0][n-1][0]
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int)int {
	if a < b {
		return a
	}
	return b
}