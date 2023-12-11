package binary_search

//Suppose an array of length n sorted in ascending order is rotated between 1 an
//d n times. For example, the array nums = [0,1,2,4,5,6,7] might become:
//
//
// [4,5,6,7,0,1,2] if it was rotated 4 times.
// [0,1,2,4,5,6,7] if it was rotated 7 times.
//
//
// Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results
//in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
//
// Given the sorted rotated array nums of unique elements, return the minimum el
//ement of this array.
//
// You must write an algorithm that runs in O(log n) time.
//
//
// Example 1:
//
//
//Input: nums = [3,4,5,1,2]
//Output: 1
//Explanation: The original array was [1,2,3,4,5] rotated 3 times.
//
//
// Example 2:
//
//
//Input: nums = [4,5,6,7,0,1,2]
//Output: 0
//Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times
//.
//
//
// Example 3:
//
//
//Input: nums = [11,13,15,17]
//Output: 11
//Explanation: The original array was [11,13,15,17] and it was rotated 4 times.
//
//

/*
	我们定义， 蓝色为， 最小值，和最小值的右边的元素。
	红色是最小值的左边元素。

	根据定义，我们知道， n-1 这个位置一定是蓝色的. (为什么？）  n-1 这个位置，要么是最小值的位置，要么是最小值右边的元素。
 */
func findMin_closed(nums []int) int {
	n := len(nums)
	left, right := 0, n-2 // [left, right]

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > nums[n-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nums[left] // l =r+1
}

/*
	用开区间写一下

	最后的结果 l+1 = r 返回 right.  (为什么？ 循环不变量是什么？ >=right 的都是蓝色, <=left 的都是红色）
 */
func findMin_open(nums []int) int {
	n := len(nums)
	left, right := -1, n-1 // (left, right)

	for left+1 < right {
		mid := (left + right) / 2
		if nums[mid] > nums[n-1] {
			left = mid
		} else {
			right = mid
		}
	}
	return nums[right] // l+1 =r
}

/*
	用半闭半开区间写一下
	这tmd 是在炫技吗？
 */
func findMin_close_open(nums []int) int {
	n := len(nums)
	left, right := 0, n-1 // [left, right)

	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[n-1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[right] // l =r
}

/*
最后来个酷的， （，] 区间
循环不变量是啥？  >right  或者说  >= right+1 的都是蓝色
 */
func findMin(nums []int) int {
	n := len(nums)
	left, right := -1, n-2 // (left, right]

	for left < right {
		mid := (left + right + 1) / 2
		if nums[mid] > nums[n-1] {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return nums[right+1] // l =r
}
