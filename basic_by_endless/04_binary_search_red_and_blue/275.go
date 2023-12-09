package binary_search

/*
太难了。做不对啊。
写出来，其它的版本。

总结，找到循环不变量。

 */
func hIndex(citations []int) int {
	n := len(citations)

	l, r := 1, n
	for l <= r {
		// 循环不变量：
		// left-1 的回答一定为「是」
		// right+1 的回答一定为「否」

		mid := l + (r-l)/2
		// 引用次数最多的 mid 篇论文，引用次数均 >= mid
		if citations[n-mid] >= mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	// 循环结束后 right 等于 left-1，回答一定为「是」
	// 根据循环不变量，right 现在是最大的回答为「是」的数
	return r
}

func hIndex2(citations []int) int {
	n := len(citations)
	l, r := 1, n+1 // [l, r)  写一个，左闭右开的区间

	for l < r {

		mid := l + (r-l)/2
		if citations[n-mid] >= mid {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r - 1
}
