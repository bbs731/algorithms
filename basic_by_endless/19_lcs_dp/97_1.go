package dp


func isInterleave(s1 string, s2 string, s3 string) bool {
	n :=len(s1)
	m := len(s2)
	if n + m != len(s3) {
		return false
	}
	var dfs func(int, int) bool

	cache := make([][]int, n)
	for i:=0; i<n; i++ {
		cache[i] = make([]int, m)
		for j :=0; j<m; j++ {
			cache[i][j]= -1
		}
	}

	dfs = func(i, j int)bool {
		if i<0 {
			return s2[:j+1] == s3[:j+1]
		}
		if j < 0 {
			return s1[:i+1] == s3[:i+1]
		}
		var ans bool
		if cache[i][j] != -1 {
			return cache[i][j] == 1
		}

		if s1[i] == s3[i+j+1]{
			 ans =  dfs(i-1, j)
		}
		if s2[j] == s3[i+j+1] {
			ans = ans || dfs(i, j-1)
		}
		if ans {
			cache[i][j] = 1
		} else{
			cache[i][j]= 0
		}
		return ans
	}
	return dfs(n-1, m-1)
}
