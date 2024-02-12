package z_algorithm


func countMatchingSubarrays(nums []int, pattern []int) int {
	m := len(pattern)
	for i:=1; i<len(nums);i++ {
		if nums[i] > nums[i-1] {
			pattern = append(pattern, 1)
		} else if nums[i] == nums[i-1]{
			pattern = append(pattern, 0)
		} else {
			pattern = append(pattern, -1)
		}
	}

	n :=len(pattern)
	z := make([]int, n)
	ans := 0
	for i, l, r :=1, 0, 0; i <n; i++ {
		if i <r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] <n &&  pattern[i+z[i]] == pattern[z[i]]{
			l, r = i, i+z[i]
			z[i]+=1
		}
		if z[i] >= m && i >=m {
			ans++
		}
	}
	return ans
}
