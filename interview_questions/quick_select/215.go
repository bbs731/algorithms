package quick_select

import "math/rand"

//func partition(nums []int, k int) int {
//	n := len(nums)
//	l, r := 0, n-1
//	rand.Shuffle(n, func(i, j int) {
//		nums[i], nums[j] = nums[j], nums[i]
//	})
//	for l < r {
//		i, j := l, r+1 // 初始化，非常的重要  用 l, r 初始化， 不是0, n-1
//		v := nums[l]
//		for {
//			for i++; i < r && nums[i] < v; i++ {
//			}
//			for j--; j > l && nums[j] > v; j-- {
//			}
//			if i >= j {
//				break
//			}
//			nums[i], nums[j] = nums[j], nums[i]
//		}
//		nums[l], nums[j] = nums[j], nums[l]
//		if j == k {
//			break
//		} else if j < k {
//			l = j + 1
//		} else {
//			r = j - 1
//		}
//	}
//	return nums[k]
//}
//

func findKthLargest(nums []int, k int) int {
	return partition(nums, len(nums)-k)
}

func partition(nums []int, k int) int {
	n := len(nums)
	l, r := 0, n-1
	rand.Shuffle(n, func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	for l < r {
		v := nums[l]
		i, j := l, r+1 // 初始化 0， n-1 bug
		for true {
			for i++; i < r && nums[i] < v; i++ { // i <= r bug
			}
			for j--; j > l && nums[j] > v; j-- { // nums[i] bug 无限循环。
			}
			if i >= j {
				break
			}
			//swap
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[l], nums[j] = nums[j], nums[l]
		if j == k {
			return nums[j]
		} else if j > k {
			r = j - 1
		} else {
			l = j + 1
		}
	}
	return nums[l] // bug  返回 -1 不对！
}
