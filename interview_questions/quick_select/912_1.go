package quick_select

import "math/rand"

func partition(nums []int, l, r int, k int) {
	for l < r {
		i, j := l, r+1
		for {
			v := nums[l]
			for i++; i < r && nums[i] < v; i++ {
			}

			for j--; j > l && nums[j] > v; j-- {
			}
			if i >= j {
				break
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[l], nums[j] = nums[j], nums[l]
		if j == k {
			break
		} else if j < k {
			l = j + 1
		} else {
			r = j - 1
		}
	}
}

func quicksort(nums []int, l, r int) {
	if l >= r {
		return
	}

	// 我kwo, 这里，是亮点把， shuffle [l, r] 区间， 不能从 0 开始！
	rand.Shuffle(r-l+1, func(i, j int) {
		nums[i+l], nums[j+l] = nums[j+l], nums[i+l]
	})

	k := (l + r) >> 1
	partition(nums, l, r, k)
	quicksort(nums, l, k-1)
	quicksort(nums, k+1, r)
}

func sortArray(nums []int) []int {
	quicksort(nums, 0, len(nums)-1)
	return nums
}
