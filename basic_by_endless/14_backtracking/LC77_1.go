package backtracking

// 这是思路2， 枚举选哪个的思路
func combine(n int, k int) (ans [][]int) {

	path := []int{}
	var dfs func(int, int)

	dfs = func(i int, d int) {
		if d == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		// 剪枝
		if i < d {
			return
		}
		// 相当于我们要求  i>=d,  既然  下面的 loop j 初始化 j=i 那么结束条件 j > 0 可以调整为  j >=d  即  j > d-1
		// 从大到小遍历
		// 可以把剪枝掉件，写在判断条件上即为 ： for j:=i; j >= d; j--  或者  j > d-1 也可以
		for j := i; j > 0; j-- {
			path = append(path, j)
			dfs(j-1, d-1)
			path = path[:len(path)-1]
		}
	}

	dfs(n, k)
	return
}
