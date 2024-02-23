package dfs

/***
N-queen 是你的强项啊！
 */

func solveNQueens(n int) (ans [][]string) {

	var dfs func(int, int, []int)

	dplus := make([]int, 2*n)
	dminus := make([]int, 2*n)

	dfs = func(i int, mask int, pos []int) {
		if i == n {
			// record the answer
			am := []string{}
			for r := 0; r < n; r++ {
				rs := make([]byte, n)
				for j := 0; j < n; j++ {
					rs[j] = '.'
				}
				rs[pos[r]] = 'Q'
				am = append(am, string(rs))
			}
			ans = append(ans, am)
			return
		}

		for j := 0; j < n; j++ {
			if mask&(1<<j) != 0 {
				continue
			}
			if dplus[i+j] != 0 || dminus[j-i+n] != 0 {
				continue
			}

			pos[i] = j
			dplus[i+j] = 1
			dminus[j-i+n] = 1
			dfs(i+1, mask|1<<j, pos)
			dplus[i+j] = 0
			dminus[j-i+n] = 0
		}
	}

	dfs(0, 0, make([]int, n))

	return ans
}
