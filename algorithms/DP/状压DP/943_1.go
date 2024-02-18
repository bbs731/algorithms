package dp

import "fmt"

func shortestSuperstring(words []string) string {
	n := len(words)
	m := 1 << uint(n)

	parent := make([][]int, m-1)
	for i := range parent {
		parent[i] = make([]int, n)
		for j := range parent[i] {
			parent[i][j] = -1
		}
	}

	prefix := preCalculation(words)

	var dfs func(int, int, int) int
	dfs = func(i int, bitmask int, overlapped int) int {
		if bitmask == m-1 {
			return overlapped
		}

		res := overlapped
		for k := range words {
			if (1<<uint(k))&bitmask != 0 {
				continue
			}
			pl := prefix[i][k]
			l := dfs(k, bitmask|1<<k, overlapped+pl)
			if l > res {
				res = l
				parent[bitmask][i] = k
			}
		}
		return res
	}

	tmp := 0
	p := 0
	for i := range words {
		l := dfs(i, 1<<i, 0)
		if l > tmp {
			tmp = l
			p = i
		}
	}
	// now start from p to build the answer
	t := 0
	perm := make([]int, n)
	seen := make([]bool, n)
	mask := 0
	for p != -1 {
		perm[t] = p
		t++
		seen[p] = true
		p2 := parent[mask][p]
		mask |= 1 << p
		p = p2
	}
	for i := 0; i < n; i++ {
		if seen[i] == false {
			perm[t] = i
			t++
		}
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
