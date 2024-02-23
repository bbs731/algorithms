package quick_select

import "sort"

/***
经典的题目啊！多看吧！
 */
func searchAposition(nums1 []int, nums2 []int, k int) int {
	m := len(nums1)
	n := len(nums2)
	h1, h2 := 0, 0

	for k > 0 {
		if h1 == m {
			return nums2[h2+k-1]
		}
		if h2 == n {
			return nums1[h1+k-1]
		}

		if k == 1 {
			return min(nums1[h1], nums2[h2])
		}

		half := k / 2

		//// 处理 index, 不要 处理 number, number 有重复， 但是 index 是唯一的。
		index1 := min(h1+half, m) - 1
		index2 := min(h2+half, n) - 1
		// 上面的index 处理的非常的巧妙， 是这道题的重点！

		//n1 := nums1[min(h1+half, m)-1]
		//n2 := nums2[min(h2+half, n)-1]
		n1 := nums1[index1]
		n2 := nums2[index2]

		if n1 <= n2 {
			k -= index1 - h1 + 1
			h1 = index1 + 1

			// 这里的二分其实是不需要的！ 可以去掉，也能算出来最终的解。 想想为什么？
			p2 := sort.SearchInts(nums2[h2:], n1)
			if p2 > 0 {
				k -= p2
				h2 = h2 + p2
			}
		} else {
			k -= index2 - h2 + 1
			h2 = index2 + 1
			// 同上面一样， 这里的二分也是可以去掉的！
			p1 := sort.SearchInts(nums1[h1:], n2)
			if p1 > 0 {
				k -= p1
				h1 = h1 + p1
			}
		}
	}

	// panic(here!) 永远到达不了这里！
	return -1
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	if (m+n)&1 == 1 {
		p := (m + n) >> 1
		return float64(searchAposition(nums1, nums2, p+1))
	}

	p1 := (n + m) >> 1
	p2 := p1 + 1
	return float64(searchAposition(nums1, nums2, p1)+searchAposition(nums1, nums2, p2)) / 2
}
