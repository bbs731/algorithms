package sort

import (
	"math/rand"
	"sort"
)

func quickSelect(a []int) int {
	k := len(a) / 2
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	// [l, r] 的闭区间
	for l, r := 0, len(a)-1; l < r; {
		v := a[l]
		i, j := l, r+1
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
	return k
}

func sortColors(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	p := quickSelect(nums)
	sortColors(nums[:p])
	sortColors(nums[p+1:])
}
