package weekly

import "sort"

/**
分数：2062
第一遍就写对了， 除了没比较右端点 right !=n

https://leetcode.cn/problems/maximum-fruits-harvested-after-at-most-k-steps/solutions/2254860/hua-dong-chuang-kou-jian-ji-xie-fa-pytho-1c2d/
灵神的代码，更简洁一些，使用了 sliding window，就不需要使用 psum 前缀和数组了。

我的逻辑，没问题，思路也很清楚，额外需要一个前缀和数组。代码稍微繁琐(有特判的情况，right ==n , right--)。

复习的时候，考虑一下，用 sliding window 重新写下题解。

 */
func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	n := len(fruits)
	values := make(map[int]int, n)
	pos := make([]int, n)

	for i, v := range fruits {
		values[v[0]] = v[1]
		pos[i] = v[0]
	}

	if _, ok := values[startPos]; !ok {
		values[startPos] = 0
		//pos = append(pos, startPos)
		mid := sort.SearchInts(pos, startPos)
		pos = append(pos[:mid], append([]int{startPos}, pos[mid:]...)...)
		//sort.Ints(pos)
		n = n + 1
	}

	psum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		psum[i] = psum[i-1] + values[pos[i-1]]
	}

	mid := sort.SearchInts(pos, startPos)
	left := sort.SearchInts(pos, startPos-k)
	right := sort.SearchInts(pos, startPos+k)

	if right == n || pos[right]-startPos > k { //right ==n  右端点容易 overflow, 犯了两次错误在这里。
		right--
	}

	ans := 0
	ans = max(ans, psum[mid+1]-psum[left])
	ans = max(ans, psum[right+1]-psum[mid])

	// 发现了，只需要是枚举，一般就不会漏答案，不要自己去拼搜索的范围。
	for i := left; i <= right; i++ {
		tmp := 0
		if i < mid {
			//if (startPos-pos[i])*2 <= k {
			tmp = psum[mid+1] - psum[i]
			rightmost := sort.SearchInts(pos, startPos+k-2*(startPos-pos[i]))
			if rightmost == n || pos[rightmost]-startPos > k-2*(startPos-pos[i]) { // check rightmost overflow
				rightmost--
			}
			if rightmost >= mid+1 {
				tmp += psum[rightmost+1] - psum[mid+1]
			}
			ans = max(ans, tmp)
			//}
		} else if i > mid {
			//if (pos[i]-startPos)*2 <= k {
			tmp = psum[i+1] - psum[mid]
			leftmost := sort.SearchInts(pos, startPos-(k-2*(pos[i]-startPos)))
			if leftmost <= mid-1 {
				tmp += psum[mid] - psum[leftmost]
			}
			ans = max(ans, tmp)
			//}
		}
	}
	return ans
}
