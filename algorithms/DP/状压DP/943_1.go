package dp

import "fmt"
func shortestSuperstring(words []string) string {
	n := len(words)
	m := 1 << uint(n)

	parent := make([][]int, m)
	dp := make([][]int, m)
	for i := range parent {
		parent[i] = make([]int, n)
		dp[i] = make([]int, n)
		for j := range parent[i] {
			parent[i][j] = -1
		}
	}
	prefix := preCalculation(words)

	for mask :=0; mask < m; mask++ {
		for bit :=0; bit <n; bit++ {
			if mask>>bit & 1 > 0 {
				pmask := mask ^( 1<<bit)
				if pmask == 0 {
					continue
				}
				for i:=0; i<n; i++ {
					if pmask >>i & 1 >  0 {
						val := dp[pmask][i] + prefix[i][bit]
						if val > dp[mask][bit]{
							dp[mask][bit] = val
							parent[mask][bit] = i
						}
					}
				}
			}
		}
	}

	//tmp := 0
	p := 0
	for j:=0; j<n; j++ {
		if dp[m-1][j] > dp[m-1][p]{
			p = j
		}
	}

	// now start from p to build the answer
	t := 0
	perm := make([]int, n)
	seen := make([]bool, n)
	mask := m-1
	for p != -1 {
		perm[t] = p
		t++
		seen[p] = true
		p2 := parent[mask][p]
		mask ^= 1 << p
		p = p2
	}

	for i := 0; i < n; i++ {
		if seen[i] == false {
			perm[t] = i
			t++
		}
	}

	for i,j :=0, t-1; i<j; i++ {
		perm[i], perm[j] = perm[j], perm[i]
		j--
	}

	fmt.Println(perm)
	// construct the answer
	ans := words[perm[0]]
	for i := 1; i < n; i++ {
		//pl := prefix[perm[i-1]][perm[i]]
		pl := prefix[perm[i-1]][perm[i]]
		ans += words[perm[i]][pl:]
	}
	return ans

}

func preCalculation(words []string) [][]int {
	n := len(words)
	prefix := make([][]int, n)
	for i := range prefix {
		prefix[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			prefix[i][j] = calculate_prefix(words[i], words[j])
		}
	}
	return prefix
}

func calculate_prefix(s1, s2 string) int {
	//l1 := len(s1)
	l2 := len(s2)
	s := s2 + s1
	n := len(s)

	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
	}

	l := 0
	for i := n - l2; i < n; i++ {
		if z[i] == n-i {
			l = n - i
			break
		}
	}
	return l
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

