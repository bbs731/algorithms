package dp

import (
	"fmt"
	"sort"
)

type Envolop struct {
	cards [][]int
}

func (e *Envolop) Len() int {
	return len(e.cards)
}

func (e *Envolop) Less(i, j int) bool {
	if e.cards[i][0] < e.cards[j][0] {
		return true
	}
	if e.cards[i][0] == e.cards[j][0] {
		return e.cards[i][1] < e.cards[j][0]
	}
	return false
}

func (e *Envolop) Swap(i, j int) {
	e.cards[i], e.cards[j] = e.cards[j], e.cards[i]
}

func maxEnvelopes(envelopes [][]int) int {
	e := &Envolop{
		envelopes,
	}
	sort.Sort(e)
	envelopes = e.cards
	fmt.Println(envelopes)
	g := [][]int{}

	for i, envelop := range envelopes {
		//pos := sort.Search(len(g), func(k int) bool { return g[k][0] >= envelopes[i][0] })
		//if pos == len(g) {
		fmt.Printf("processing: %v", envelop)
		if len(g) == 0 {
			e := make([]int, 2)
			copy(e, envelopes[i])
			g = append(g, e)
		} else {
			// 跟新，insert 或者抛弃不用

			// insertion
			if envelop[0] > g[len(g)-1][0] && envelop[1] > g[len(g)-1][1] {
				e := make([]int, 2)
				copy(e, envelop)
				g = append(g, e)
				fmt.Printf("insertion : %v\n", e)
				continue
			}

			//  抛弃不用
			if envelop[0] == g[len(g)-1][0] && envelop[1] >= g[len(g)-1][1] {
				fmt.Printf("discard: %v\n", envelop)
				continue
			}

			// 这一步是错误的， replace 很危险。		[[3 4] [12 2] [12 15] [30 50]]
			// now the case envelop[1] < g[len(g)-1][1], we need to replace
			if envelop[1] < g[len(g)-1][1] && (len(g)-2 < 0 || g[len(g)-2][1] < envelop[1]) {
				// replace
				fmt.Printf("replace: %v with %v\n", g[len(g)-1], envelop)
				g[len(g)-1][0] = envelop[0]
				g[len(g)-1][1] = envelop[1]
			}
		}
	}
	fmt.Println(g)
	return len(g)
}

// 这个solution 会超时， 需要优化 DP
func maxEnvelopes_dp(envelopes [][]int) int {
	n := len(envelopes)
	f := make([]int, n)
	ans := 0

	e := &Envolop{
		envelopes,
	}
	sort.Sort(e)
	envelopes = e.cards

	for i := range envelopes {
		local := 1
		for j := 0; j < i; j++ {
			if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
				local = max(local, f[j]+1)
				break
			}
		}
		f[i] = local
		ans = max(ans, local)
	}
	return ans
}

