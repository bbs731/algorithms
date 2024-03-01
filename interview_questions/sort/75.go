package sort

func partition(nums []int, l, r int, k int) {
	if l >= r {
		return
	}

	//n := len(nums)
	//l, r := 0, n-1

	for l < r {
		v := nums[l]
		// 这里容易出现巨大的bug, i 不能初始化成 l-1 想一下为什么啊？
		// 确实是危险的， i 初始位 l-1， 就可能把 nums[l] 这个位置 swap 出去了， nums[l] 这个位置是 pivot 不能动的！
		// 找这个 bug 花了半个小时啊！
		i, j := l-1, r+1
		for true {
			for i++; i < r && nums[i] < v; i++ {
			}
			for j--; j > l && nums[j] > v; j-- {
			}
			if i >= j {
				break
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[l], nums[j] = nums[j], nums[l]
		if j == k {
			break
		} else if j > k {
			r = j - 1
		} else {
			l = j + 1
		}
	}
	return
}

func sortColors(nums []int) {
	n := len(nums)
	if n <= 1 { // 这里的判断，也是需要的， 要不然有可能会出现  infinite recusion loop
		return
	}

	mid := n >> 1
	partition(nums,0, n-1, mid)
	sortColors(nums[:mid])
	sortColors(nums[mid+1:])
}
