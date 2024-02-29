package sort

import "math/rand"

func quickSelect(nums []int, k int) {
	n := len(nums)
	rand.Shuffle(n, func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	l, r := 0, n-1
	for l < r {
		v := nums[l]
		i, j := l, r+1  // 这里容易出现 bug,  i 不能初始化成 l-1
		for {
			for i++; i<r && nums[i]< v; i++  {}
			for j--; j>l && nums[j]>v; j--{}
			if i >=j {
				break
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[l], nums[j]= nums[j], nums[l]
		if j == k{
			break
		} else if j > k {
			r = j-1
		} else {
			l = j+1
		}
	}
	return
}

func sortColors(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	p := n >> 1
	quickSelect(nums, p)
	sortColors(nums[:p])
	sortColors(nums[p+1:])
}
