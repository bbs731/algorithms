package quick_select

import "math/rand"

/*
leetcode quick-slect 的 template code
https://leetcode.cn/problems/kth-largest-element-in-an-array/solutions/307351/shu-zu-zhong-de-di-kge-zui-da-yuan-su-by-leetcode-/
 */

/***
https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/weekly/262/b/b.go#L38
灵神的版本， 感觉不错。

灵神写的这个，最容易记住。
就记住这个模版就可以了。

如果 l, r 的区间是 [0, n-1] 那么 k 的取值也是 [0, n-1] 之间， 就是说 k 的 index 是从 0 开始的。
k 的 index 从 0 开始

 */
func quickSelect(a []int, l, r, k int) int {
	//k := len(a) / 2
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	// [l, r] 的闭区间
	// l , r := 0, len(a) -1
	for l < r {
		v := a[l]
		i, j := l, r+1 // 这个是灵神的代码， i 初始化成 l 是对的。
		for {
			for i++; i < r && a[i] < v; i++ {
			}
			for j--; j > l && a[j] > v; j-- {
			}
			if i >= j {
				break
			}
			a[i], a[j] = a[j], a[i]
		}
		// j 和 l d的数交换一下
		a[l], a[j] = a[j], v
		if j == k {
			break
		} else if j < k {
			l = j + 1
		} else {
			r = j - 1
		}
	}
	return a[k]
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
