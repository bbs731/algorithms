package binary_search

import "sort"

/****

给你一个 n x n 矩阵 matrix ，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
请注意，它是 排序后 的第 k 小元素，而不是第 k 个 不同 的元素。

你必须找到一个内存复杂度优于 O(n^2) 的解决方案。



示例 1：

输入：matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
输出：13
解释：矩阵中的元素为 [1,5,9,10,11,12,13,13,15]，第 8 小元素是 13
示例 2：

输入：matrix = [[-5]], k = 1
输出：-5


提示：

n == matrix.length
n == matrix[i].length
1 <= n <= 300
-10^9 <= matrix[i][j] <= 10^9
题目数据 保证 matrix 中的所有行和列都按 非递减顺序 排列
1 <= k <= n^2

 */

func kthSmallest(matrix [][]int, k int) int {
	// 我们来二分一下答案
	l, r := -int(1e9)-1, int(1e9)
	n := len(matrix)

	// 先 false 后 true 的 正常序列。
	// 因为可能有重复元素， 是不是使用 （，] 左开右闭的区间会更好些？
	for l < r {
		mid := (l + r + 1) >> 1

		rank := 0
		for i := 0; i < n; i++ {
			if mid >= matrix[i][n-1] {
				rank += n
			} else {
				p := sort.SearchInts(matrix[i], mid+1) - 1
				rank += p + 1
			}
		}
		if rank >= k {
			r = mid - 1
		} else {
			l = mid
		}
	}
	// l == r
	return r + 1
}

/*****

wow!  太牛了， 每次都能提供双解。

这是一个，正常 先 false, 后 true 的序列， 所以不需要取反 ++ 等逻辑。

只需要根据值域，来做 shift 就可以
sort.Search() 的使用技巧二
 */

func kthSmallest(matrix [][]int, k int) int {
	// 我们来二分一下答案
	//l, r := -int(1e9)-1, int(1e9)
	n := len(matrix)

	// [l, r)
	return -int(1e9) + sort.Search(2*int(1e9)+1, func(mid int) bool {
		mid += -int(1e9)
		rank := 0
		for i := 0; i < n; i++ {
			if mid >= matrix[i][n-1] {
				rank += n
			} else {
				p := sort.SearchInts(matrix[i], mid+1) - 1
				rank += p + 1
			}
		}
		return rank >= k
	})
}

/***
	灵神的答案， 比我写的好： 上下界，写的比我的窄。
 */

func kthSmallest(a [][]int, k int) int {
	// 注意 k 从 1 开始
	n, m := len(a), len(a[0])
	mi, mx := a[0][0], a[n-1][m-1]
	ans := sort.Search(mx-mi, func(v int) bool {
		v += mi
		cnt := 0
		for i, j := 0, m-1; i < n && j >= 0; {
			if v < a[i][j] {
				j--
			} else {
				cnt += j + 1
				i++
			}
		}
		return cnt >= k
	}) + mi
	return ans
}
