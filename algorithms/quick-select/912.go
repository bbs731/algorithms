package quick_select

import "math/rand"

// k 的 index 从 0 开始
func partition(nums []int, l, r, k int) {
	rand.Shuffle(r-l+1, func(i, j int) {
		nums[l+i], nums[l+j] = nums[l+j], nums[l+i]
	})
	for l < r {
		i, j := l, r+1
		v := nums[l]
		for {
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
			return
		} else if j > k {
			r = j - 1
		} else {
			l = j + 1
		}
	}
}

func quicksort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	partition(nums, l, r, mid)
	quicksort(nums, l, mid-1)
	quicksort(nums, mid+1, r)
}

func sortArray(nums []int) []int {
	quicksort(nums, 0, len(nums)-1)
	return nums
}
