package quick_select

import "math/rand"

// find the kth element, k  的 index 从 0 开始
func partition(nums []int, k int) int {
	n := len(nums)
	l, r := 0, len(nums)-1
	rand.Shuffle(n, func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })

	for l < r {
		i, j := 0, r+1
		v := nums[l]
		for {
			// 这里的条件，好像也很重要啊！
			for i++; i < r && nums[i] < v; i++ {
			}
			for j--; j > l && nums[j] > v; j-- {
			}

			if i >= j {
				break
			}
			// swap
			nums[i], nums[j] = nums[j], nums[i]
		}

		// swap l, and j
		nums[l], nums[j] = nums[j], nums[l]
		if j == k {
			break
		} else if j < k {
			// 只涉及到 index 的计算， 不涉及到， k 的加减， 方便了许多啊！
			l = j + 1
		} else {
			r = j - 1
		}
	}

	return nums[k]
}
