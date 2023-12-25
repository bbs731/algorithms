package quick_select

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickselect_loop(nums, 0, n-1, n-k)
}

// 写成左闭右开是对的， 为啥写成闭区间是错的？
// recursive to loop 怎么选择区间？

func quickselect_loop(nums []int, l, r, k int) int {
	for l < r { // [l, r)  左闭右开区间
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
			r = j
		} else {
			//return quickselect(nums, j+1, r, k)
			l = j + 1
		}
	}
	return nums[l] // l == r，  return nums[r] 也是可以的。
}
