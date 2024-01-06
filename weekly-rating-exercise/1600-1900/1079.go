package _600_1900



// 看看 灵神的题解， O（n^2) 的解法
//https://leetcode.cn/problems/letter-tile-possibilities/solutions/2275356/on2-ji-shu-dppythonjavacgo-by-endlessche-hmez/
//状态转移方程，很难定义啊！

// 这是好题，将来重新做！




// dfs 过，是因为 len(tiles) <=7  再大就过不去了。
func numTilePossibilities(tiles string) int {
	n := len(tiles)
	ans := 0
	tables := make(map[string]bool)
	tables[""] = true

	var dfs func(int, string)
	dfs = func(i int, content string){
		if i == n {
			if _, ok := tables[content]; !ok {
				ans++
				tables[content]= true
			}
			return
		}

		c := string(tiles[i])
		// 不选
		dfs(i+1, content)
		dfs(i+1, c+content)
		dfs(i+1, content+c)
		for j :=1; j<len(content); j++ {
			dfs(i+1, content[:j] + string(tiles[i]) + content[j:])
		}
	}
	dfs(0, "")
	return ans
}