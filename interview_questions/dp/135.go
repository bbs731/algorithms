package dp

import "fmt"

/***

n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

你需要按照以下要求，给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。



示例 1：

输入：ratings = [1,0,2]
输出：5
解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
示例 2：

输入：ratings = [1,2,2]
输出：4
解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。

 */

/****

像是 sliding window 的问题。

这种题，好像是不难的， 但是需要 reasonning

 */
func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 || n == 1 {
		return n
	}

	f := make([]int, n)
	for i := range f {
		f[i] = 1
	}

	for i := 0; i < n; {
		j := i + 1
		for ; j < n && ratings[j-1] > ratings[j]; j++ {
		}
		// 分类讨论一下
		f[i] = max(f[i], j-i)
		for cnt, k := 1, j-1; k >= i+1; k-- {
			f[k] = cnt
			cnt++
		}

		if i > 0 && ratings[i] > ratings[i-1] {
			f[i] = max(f[i], f[i-1]+1)
		}
		i = j
	}
	// 对啊， 感觉这里就不需要特殊的处理。
	//// 需要处理 n-1 这个 element, 为什么呢？
	//if ratings[n-1] > ratings[n-2] {
	//	f[n-1] = max(f[n-1], f[n-2]+1)
	//}

	ans := 0
	for i := 0; i < n; i++ {
		ans += f[i]
	}
	fmt.Println(f)
	return ans
}

/***
做一下 空间的优化， 这道题，统计 ans, 没必要保存 f 数组， 变计算，边统计就可以。
优化的时候，有风险啊， 面试的时候慎用啊！
 */
func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 || n == 1 {
		return n
	}
	ans := 0
	prev := 1
	for i := 0; i < n; {
		current := 1
		j := i + 1
		for ; j < n && ratings[j-1] > ratings[j]; j++ {
		}
		// 分类讨论一下
		prevc := 1
		for cnt, k := 1, j-1; k >= i+1; k-- {
			if k == j-1 {
				prevc = cnt
			}
			ans += cnt
			cnt++
		}

		current = max(j-i, current)
		if i > 0 && ratings[i] > ratings[i-1] {
			current = max(current, prev+1)
		}
		if i+1 == j {
			prev = current
		} else {
			prev = prevc
		}
		ans += current
		fmt.Println(current)
		i = j
	}

	return ans
}
