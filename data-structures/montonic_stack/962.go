package montonic_stack

import (
	"sort"
)

/***
惩罚一下，你第一次写， 写了1个小时， binary_search 全弄错。
现在把所有的写法，都写一遍！

很好的模版题目。 这要是面试，第一次遇到这题，又死了！
 */
func maxWidthRamp(nums []int) int {
	n := len(nums)
	st := make([]int, 0)
	heights := make([]int, 0)
	ans := 0

	for i:=n-1; i>=0; i-- {
		if len(st)==0 || nums[i] > nums[st[len(st)-1]] {
			st = append(st, i)
			heights = append(heights, nums[i])
			continue
		}
		// 我的天啊， 为了写这个 search 语句， 我写了一个小时。
	 	p := sort.SearchInts(heights, nums[i])
	 	if p != len(heights) {
			//fmt.Println(nums[i], heights[p])
			ans = max(ans, st[p]-i)
		}
	}
	return ans
}

func maxWidthRamp(nums []int) int {
	n := len(nums)
	st := make([]int, 0)
	//heights := make([]int, 0)
	ans := 0

	for i:=n-1; i>=0; i-- {
		if len(st)==0 || nums[i] > nums[st[len(st)-1]] {
			st = append(st, i)
			//heights = append(heights, nums[i])
			continue
		}
		// 我的天啊， 为了写这个 search 语句， 我写了一个小时。
		//p := sort.SearchInts(heights, nums[i])
		p := sort.Search(len(st), func (mid int) bool {
			return nums[st[mid]] >= nums[i]
		})
		if p != len(st) {
			ans = max(ans, st[p]-i)
		}
	}
	return ans
}




func maxWidthRamp(nums []int) int {
	n := len(nums)
	st := make([]int, 0)
	//heights := make([]int, 0)
	ans := 0

	for i:=n-1; i>=0; i-- {
		if len(st)==0 || nums[i] > nums[st[len(st)-1]] {
			st = append(st, i)
			//heights = append(heights, nums[i])
			continue
		}
		// 我的天啊， 为了写这个 search 语句， 我写了一个小时。
		//p := sort.SearchInts(heights, nums[i])
		m := len(st)
		l, r := -1, m
		for l + 1 < r {
			mid := (l+r)>> 1
			if nums[st[mid]] < nums[i]{  // 这里, 写错了，浪费了1个小时， target 写在后面。
				l = mid
			} else {
				r = mid
			}
		}
		p := r
		if p != len(st) {
			ans = max(ans, st[p]-i)
		}
	}
	return ans
}


func maxWidthRamp(nums []int) int {
	n := len(nums)
	st := make([]int, 0)
	//heights := make([]int, 0)
	ans := 0

	for i:=n-1; i>=0; i-- {
		if len(st)==0 || nums[i] > nums[st[len(st)-1]] {
			st = append(st, i)
			//heights = append(heights, nums[i])
			continue
		}
		// 我的天啊， 为了写这个 search 语句， 我写了一个小时。
		//p := sort.SearchInts(heights, nums[i])
		m := len(st)
		l, r := -1, m-1
		for l < r {
			mid := (l+r+1)>> 1
			if nums[st[mid]] < nums[i]{  // 这里, 写错了，浪费了1个小时， target 写在后面。
				l = mid
			} else {
				r = mid-1
			}
		}
		p := r+1
		if p != len(st) {
			ans = max(ans, st[p]-i)
		}
	}
	return ans
}

func maxWidthRamp(nums []int) int {
	n := len(nums)
	st := make([]int, 0)
	//heights := make([]int, 0)
	ans := 0

	for i:=n-1; i>=0; i-- {
		if len(st)==0 || nums[i] > nums[st[len(st)-1]] {
			st = append(st, i)
			//heights = append(heights, nums[i])
			continue
		}
		// 我的天啊， 为了写这个 search 语句， 我写了一个小时。
		//p := sort.SearchInts(heights, nums[i])
		m := len(st)
		l, r := 0, m-1
		for l <= r {
			mid := (l+r)>> 1
			if nums[st[mid]] < nums[i]{  // 这里, 写错了，浪费了1个小时， target 写在后面。
				l = mid+1
			} else {
				r = mid-1
			}
		}
		// r = l -1
		p := l
		if p != len(st) {
			ans = max(ans, st[p]-i)
		}
	}
	return ans
}



func maxWidthRamp(nums []int) int {
	n := len(nums)
	st := make([]int, 0)
	//heights := make([]int, 0)
	ans := 0

	for i:=n-1; i>=0; i-- {
		if len(st)==0 || nums[i] > nums[st[len(st)-1]] {
			st = append(st, i)
			//heights = append(heights, nums[i])
			continue
		}
		// 我的天啊， 为了写这个 search 语句， 我写了一个小时。
		//p := sort.SearchInts(heights, nums[i])
		m := len(st)
		// [l, r)
		l, r := 0, m
		for l < r {
			mid := (l+r)>> 1
			if nums[st[mid]] < nums[i]{  // 这里, 写错了，浪费了1个小时， target 写在后面。
				l = mid+1
			} else {
				r = mid
			}
		}
		// l == r
		p := r
		if p != len(st) {
			ans = max(ans, st[p]-i)
		}
	}
	return ans
}
