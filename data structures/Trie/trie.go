package Trie

/*
数组实现的 trie 版本
详情见：  https://oi-wiki.org/string/trie/
 */
var next [100000][26]int
var exit [100000]bool
var cnt int

func insert(s string) {
	l := len(s)
	p := 0
	for i := 0; i < l; i++ {
		c := s[i] - 'a'
		if next[p][c] == 0 {
			cnt++
			next[p][c] = cnt
		}
		p = next[p][c]
	}
	exit[p] = true
}

func find(s string) bool {
	p := 0
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		if next[p][c] == 0 {
			return false
		}
		p = next[p][c]
	}
	return exit[p]
}
