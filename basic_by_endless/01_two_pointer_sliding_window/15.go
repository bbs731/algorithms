package sliding_window

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)

	ans := make([][]int, 0)

	for i :=range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := n - 1
		for left < right {
			if nums[left]+nums[right] < -nums[i] {
				left++
			} else if nums[left]+nums[right] > -nums[i] {
				right--
			} else {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})

				// 下面的处理太巧妙了，做不到啊！
				left++
				for left < n && nums[left] == nums[left-1] {
					left++
				}
				right--
				for right >=0 && right +1 <n && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return ans
}
