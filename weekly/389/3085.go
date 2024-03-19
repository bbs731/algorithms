package weekly

import (
	"math"
	"sort"
)

/****
https://leetcode.cn/problems/minimum-deletions-to-make-string-k-special/solutions/2692077/kao-lu-zui-duo-bao-liu-duo-shao-ge-zi-mu-qttz/

灵神的解，比较牛叉！
换答案， 哈哈！
 */
func minimumDeletions(word string, k int) int {
	freqm := make(map[int32]int)

	for _, c := range word {
		freqm[c]++
	}
	l := make([]int, 0, len(freqm))
	for _, v := range freqm {
		l = append(l, v)
	}
	sort.Ints(l)
	n := len(l)
	presum := make([]int, n+1)

	for i := 1; i <= n; i++ {
		presum[i] = presum[i-1] + l[i-1]
	}

	ans := math.MaxInt32
	// 枚举左端点？
	//fmt.Println(l)
	for i := 0; i < n; i++ {
		res := presum[i]
		//pos := sort.SearchInts(l, l[i]+k)
		//target := l[i] + k
		left, right := -1, n-1
		// need upper_bound ?
		for left < right {
			mid := (right + left + 1) >> 1
			if l[mid] <= l[i]+k {
				left = mid
			} else {
				right = mid - 1
			}
		}
		pos := right + 1

		if pos < n {
			res += l[pos] - l[i] - k
			res += presum[n] - presum[pos+1] - (l[i]+k)*(n-1-pos) // 我靠， 好不容易才算对啊！
		}
		ans = min(ans, res)
	}
	return ans
}
