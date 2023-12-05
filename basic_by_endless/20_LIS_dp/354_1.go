package dp

import (
	"fmt"
	"sort"
)


/*
https://leetcode.cn/problems/russian-doll-envelopes/solutions/633231/e-luo-si-tao-wa-xin-feng-wen-ti-by-leetc-wj68/

官方的题解说的很清楚！ 关键解释了， 为什么给 envolop 排序的时候，第二个维度要是降序的， 这个太赞了！
然后就姜维到了，一维的 LIS 问题用  n *logn 复杂度解决。


 */
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
		return e.cards[i][1] > e.cards[j][1]  // 这里为什么要倒着排列呢？
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
	g := []int{}

	for _, envelop := range envelopes {
		pos := sort.SearchInts(g, envelop[1] )
		if pos == len(g){
			g = append(g, envelop[1])
		} else {
			g[pos] = envelop[1]
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
			}
		}
		f[i] = local
		ans = max(ans, local)
	}
	return ans
}

