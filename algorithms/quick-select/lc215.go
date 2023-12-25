package quick_select

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickselect_loop(nums, 0, n-1, n-k)
}

// 写成左闭右开是对的， 为啥写成闭区间是错的？  我感觉就是碰巧了， 在双指针的 loop 里 有 j := r +1  只适合 [ ) 左闭右开区间
// 为什么 闭区间不行， 把 partition 写成标准的样子，闭区间写是没有问题的， 但是，partition 里面用了双指针的技巧，造成，写闭区间不行。
// 因为有 i := l - 1 和 j := r+ 1 最后造成返回下标，有的时候 返回i 正确，有的时候返回 j 正确，无法判断。 但是标准的 partition 版本，
// 只move 一个指针，写闭区间，正确性是没问题的，但是这道 215的题会超时。
// recursive to loop 怎么选择区间？ 这道题是碰巧不要深究。

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
			r = j // loop 里正好有一个  j := r+1  所以天然的凑巧适合了左闭右开区间
		} else {
			//return quickselect(nums, j+1, r, k)
			l = j + 1
		}
	}
	return nums[l] // l == r，  return nums[r] 也是可以的。
}
