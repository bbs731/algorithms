package interview_questions

import "sort"

/***
给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。

当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

注意：不允许旋转信封。


示例 1：

输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出：3
解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
示例 2：

输入：envelopes = [[1,1],[1,1],[1,1]]
输出：1


提示：

1 <= envelopes.length <= 105
envelopes[i].length == 2
1 <= wi, hi <= 105
 */

/***
这 tmd 是在作弊啊！
 */

func maxEnvelopes(envelopes [][]int) int {
	//n := len(envelopes)
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] > b[1])
	})

	g := []int{}
	for _, e := range envelopes {
		p := sort.SearchInts(g, e[1])
		if p == len(g) {
			g = append(g, e[1])
		} else {
			g[p] = e[1]
		}
	}
	return len(g)
}
