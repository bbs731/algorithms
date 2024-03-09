package interview_questions

import "sort"

/***
集合去重， 有什么手段吗?
 */

func combinationSum2(candidates []int, target int) [][]int {
	// 利用枚举 freq 来去重
	freq := make(map[int]int)
	for _, num := range candidates {
		freq[num]++
	}

	nums := make([]int, 0, len(freq))
	for k := range freq {
		nums = append(nums, k)
	}
	n := len(nums)

	ans := make([][]int, 0)
	sort.Ints(nums)
	var dfs func(int, int, []int)
	dfs = func(i, s int, l []int) {
		if i == n {
			if s == target {
				ans = append(ans, append([]int{}, l...))
			}
			return
		}

		dfs(i+1, s, l)
		for k := 1; k <= freq[nums[i]]; k++ {
			if s+k*nums[i] <= target {
				for p := 1; p <= k; p++ {
					l = append(l, nums[i])
				}
				dfs(i+1, s+k*nums[i], l)
				l = l[:len(l)-k]
			}
		}
	}
	dfs(0, 0, []int{})
	return ans
}
