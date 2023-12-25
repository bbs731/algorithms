package quick_select

/*
leetcode quick-slect 的 template code
https://leetcode.cn/problems/kth-largest-element-in-an-array/solutions/307351/shu-zu-zhong-de-di-kge-zui-da-yuan-su-by-leetcode-/
 */

// 利用双指针，避免最坏的 O(n^2) 的情况， 可以做到平均时间复杂度是 O(n)
// 这段代码写的太华丽了！  这里的 kth index 从 0 开始算， 第 3th element 实际上是找排序之后，对应的下标 3 （实际上是第 4小的 element)
// quick-select 算法的前提是 nums 数组是可以 in-place 改动的。
func quickselect(nums []int, l, r, k int) int {
	if l == r {
		return nums[l]
	}
	pivot := nums[l]
	i := l - 1
	j := r + 1
	//使用双指针的方法，这种方法能够较好地应对各种数据。
	for i < j {
		// 尤其是这连个循环，写得太华丽了！
		for i++; nums[i] < pivot; i++ {
		}
		for j--; nums[j] > pivot; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickselect(nums, l, j, k)
	}
	return quickselect(nums, j+1, r, k)
}

// quick-select  也可以写成下面分成两步的
/*
1. partition 函数
2. KthSmallestArray recursive 函数
 */

func partition(nums []int, l, r int) int {
	pivot := nums[l]
	i := l - 1
	j := r + 1
	//使用双指针的方法，这种方法能够较好地应对各种数据。
	for i < j {
		// 尤其是这连个循环，写得太华丽了！
		for i++; nums[i] < pivot; i++ {
		}
		for j--; nums[j] > pivot; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return j // 这里 很重要一定要返回 j 不能是 i 为什么呢？ 因为结束的条件有可能是 j+1= i  也有可能是 i=j 所以应该返回 j 而不是 i
	// 所以最后得到的循环不变量是  <=j 的 elements 的值都 <= pivot
}

func KthSmallestArray(nums []int, l, r, k int) int {
	if l == r {
		return nums[l]
	}
	pos := partition(nums, l, r)

	if k <= pos {
		return KthSmallestArray(nums, l, pos, k)
	}
	return KthSmallestArray(nums, pos+1, r, k)
}

///下面 这个 partition 和 kth samlllest array 的写法会超时
//忘记下面垃圾的partition的 写法， 记住用双指针来写 partition
// abondoned
func partition(nums []int, l, r int) int {
	pivot := nums[r] // 这个 pivot 的选取依赖 nums 是随机的，对于精心策划的数据，有可能表现不佳!
	i := l - 1       // 类似 quick sort partition 的发方式

	for j := l; j < r; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i] // swap i and j
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}
