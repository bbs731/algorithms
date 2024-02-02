package dp

import "sort"

/***

这是一套， 可以处理， 线性DP 含有 intervals 的通用模板。 （改进了，之前自己更长的模板）
可以通用的解决，多个 intervals 有共同 end time 的问题， 不依赖于 end time 的值域。（过程中做了，end time 的离散化）
可以解决 template question: 1235, 2830, 2054 等问题。

时间复杂度是 O(n*logn)

需要注意的是， 虽然是个通用的模板，处理 Intervals 线性DP 问题， 但是有的问题是可以在 O(n) 的时间复杂度就能解决掉了，
例如 2830, 2008 等问题， 所以，对于简单的问题，还是应该看出来，可以一次循环搞定的，不需要 sort and binary search

 */

func dp_binary_search_template_code(offers [][]int) int {
	n := len(offers)
	type record struct {
		start, end, value int
	}
	records := make([]record, n+1)
	for i := 0; i < n; i++ {
		start, end, value := offers[i][0], offers[i][1], offers[i][2]
		records[i] = record{start, end, value}
	}
	// dummy node
	records[n] = record{0, 0, 0}

	sort.Slice(records, func(i, j int) bool {
		return records[i].end < records[j].end
	})

	dp := make([]int, n+1) // 这里的 dp 相当于把 end 的值域离散化了，当 end 的值域是 1e9 也没关系，只对 end 出现过的次数哟关系 n
	for i, r := range records {
		if i == 0 {
			continue // 或者让 index i 从 1 开始也可以。
		}
		// 需要找打第一个  end  <= start  的 end   >= (start+1) - 1 的位置。
		j := sort.Search(i, func(j int) bool { return records[j].end >= (r.start + 1) }) - 1
		dp[i] = max(dp[i-1], dp[j]+r.value)
	}
	return dp[n]

}
