package quick_select

import "math/rand"

/***
见了鬼了， 一次写不对，就废了啊。
 */

func partition(nums []int, k int) {
	l, r := 0, len(nums)-1
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
		// swap l, and j
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

func qs(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// 这个 shuffle 至关重要！
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	mid := len(nums) >> 1
	partition(nums, mid)
	qs(nums[:mid])
	qs(nums[mid+1:])
}

func sortArray(nums []int) []int {
	qs(nums)
	return nums
}
