package weekly

func prefix_function(pattern []int) []int {
	n := len(pattern)
	pi := make([]int, n)

	j := 0 // j 记录的是 pi[i-1], 初始化为 pi[0]  即为 0
	for i := 1; i < n; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
		}

		if pattern[i] == pattern[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}

func countMatchingSubarrays(nums []int, pattern []int) int {

	pi := prefix_function(pattern)
	pos := make([]int, 0)

	n := len(nums)
	// 这个数组 b 还需要灵神帮你讲一遍，哎！
	b := make([]int, n-1)
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			b[i-1] = 1
		} else if nums[i] == nums[i-1] {
			b[i-1] = 0
		} else {
			b[i-1] = -1
		}
	}

	text := b
	j := 0
	for i := 0; i < len(text); i++ {
		for j > 0 && pattern[j] != text[i] {
			j = pi[j-1]
		}

		if pattern[j] == text[i] {
			j++
		}
		if j == len(pattern) {
			pos = append(pos, i-len(pattern)+1)
			j = pi[j-1]
		}
	}
	return len(pos)
}
