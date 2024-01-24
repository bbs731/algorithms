package main

import (
	"fmt"
	"sort"
)

/*
去看一下 upper_bound 的定义：
https://stackoverflow.com/questions/28389065/difference-between-basic-binary-search-for-upper-bound-and-lower-bound

举一个例子：
nums := []int{1, 2, 4, 4, 4, 7, 8, 9, 9}
upper_bound 查找数字 4 的 upper bound 的位置， 是 7 数字的这个位置， 不是最后一个 4的位置。

upper bound 用标准库里的函数实现：
1. sort.SearchInts(nums, 4+1)
2. sort.Search(len(nums), func (i int) bool {return nums[i] > 4})
 */

// 开区间的写法
func upper_bound(nums []int, target int, left, right int) int {
	l := left - 1
	r := right + 1

	for l+1 < r {
		mid := (r-l)/2 + l
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid
		}
	}
	// l + 1 = r
	return r
}

// 闭区间的写法 [l, r]
func upper_bound2(nums []int, target int, l, r int) int {
	//n := len(nums)
	//l, r := 0, n-1 // [l, r]

	// 还是要用红蓝染色，来考虑这个问题。
	// r+1 指向的是蓝色
	// l-1 指向的是红色
	for l <= r {
		mid := (r-l)>>1 + l
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l // l = r+1
}

func upper_bound3(nums []int, target int, left, right int) int {
	l := left
	r := right + 1

	for l < r {
		mid := (r-l)>>1 + l
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	// l == r
	return r
}

func upper_bound4(nums []int, target int, left, right int) int {
	l := left - 1
	r := right

	for l < r {
		mid := (l + r + 1) >> 1
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	// l == r
	return l + 1
}

// upper_bound 5, 6, 7 在指定， left 和 right 的搜索范围是受限的， 不如在调用之前，先对 nums 做裁剪
func upper_bound5(nums []int, target int, l, r int) int {
	//  搜索   >= target + 1  等价于   > target
	return sort.SearchInts(nums, target+1)
}

func upper_bound6(nums []int, target int, l, r int) int {
	return sort.Search(len(nums), func(i int) bool { return nums[i] > target })
}

// 这个和  upper_bound5 是一样的。
func upper_bound7(nums []int, target int, l, r int) int {
	return lower_bound(nums, target+1)
}

func lower_bound(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1 // [l, r]  闭区间
	/* 循环不变量是： l-1 指向的是红色   <=l-1 都是红色
	 				r+1 指向的是蓝色   >= r+1 都是蓝色
	*/
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l // l = r+1
}

func main() {

	num := -1
	nums := []int{1, 1, 2, 4, 5, 5, 5, 7, 7, 8, 9}
	n := len(nums)

	type f func([]int, int, int, int) int
	fl := []f{upper_bound, upper_bound2, upper_bound3, upper_bound4, upper_bound5, upper_bound6, upper_bound7}

	for _, fc := range fl {
		fmt.Println(fc(nums, num, 0, n-1))
	}
}
