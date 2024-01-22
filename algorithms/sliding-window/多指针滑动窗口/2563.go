package sliding_window

import (
	"sort"
)

/***

[0, 1, 4, 4, 5, 7]

 */

/***
这道题，做了3个小时吗？ 怎么办？

lower_bound,
upper_bound 还是心魔啊！ upper_bound 的定义， 一值模棱两可， 你不是很确定， upper_bound 指的是那个位置。

https://stackoverflow.com/questions/28389065/difference-between-basic-binary-search-for-upper-bound-and-lower-bound
 */

// 春雷你对 upper_bound 的理解，一直是有错误的。
// https://stackoverflow.com/questions/28389065/difference-between-basic-binary-search-for-upper-bound-and-lower-bound

// 找机会把 upper bound 的 几重实现， 再多写几遍！
func upper_bound(nums []int, v int, l, r int) int {
	left := l - 1
	right := r + 1

	for left+1 < right {
		mid := (left + right) / 2
		if nums[mid] <= v {
			left = mid
		} else {
			right = mid
		}
	}
	// is this upper bound?
	//return left
	return right
}

func lower_bound(nums []int, v int, l, r int) int {
	left := l - 1
	right := r + 1

	for left+1 < right {
		mid := (left + right) / 2
		if nums[mid] < v {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}

// 以下两种做法都是对的！

// 这个解法是，枚举右端点， 从右往左看。
func countFairPairs(nums []int, lower int, upper int) int64 {
	n := len(nums)
	sort.Ints(nums)

	ans := 0
	for right := 0; right < n; right++ {
		v := nums[right]

		pos_right := upper_bound(nums, upper-v, 0, right-1)
		//pos_right:= sort.SearchInts(nums[:right], upper-v+1)
		pos_left := lower_bound(nums, lower-v, 0, right-1)
		ans += pos_right - pos_left // pos_right-1 - pos_left +1  = pos_right - pos_left
	}
	return int64(ans)
}

// 这个解法，枚举左端点，从左往右看.
func countFairPairs(nums []int, lower int, upper int) int64 {
	n := len(nums)
	sort.Ints(nums)

	ans := 0
	for left := 0; left < n-1; left++ {
		v := nums[left]

		pos_right := upper_bound(nums, upper-v, left+1, n-1)
		pos_left := lower_bound(nums, lower-v, left+1, n-1)

		if pos_right >= pos_left {
			ans += pos_right - pos_left
		}
	}
	return int64(ans)
}

// 下面是暴力的做法，时间复杂度通不过。
func countFairPairs(nums []int, lower int, upper int) int64 {
	n := len(nums)
	sort.Ints(nums)

	ans := 0
	for left := 0; left < n-1; left++ {
		v := nums[left]

		for k := n - 1; k > left; k-- {
			if nums[k]+v <= upper && nums[k]+v >= lower {
				ans++
			}
		}
	}

	return int64(ans)
}
