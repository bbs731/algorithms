package binary_search

/**
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。



示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

 */

//func searchAposition(nums1, nums2 []int, pos int) int {
//	l, r := -int(1e6)-1, int(1e6)+1
//	var p1, p2 int
//	for l+1 < r {
//		mid := (l + r) >> 1
//		p1 = sort.SearchInts(nums1, mid) // 有重复的元素， 怎么办？
//		p2 = sort.SearchInts(nums2, mid)
//
//		if p1+p2 > pos-1 {
//			r = mid
//		} else if p1+p2 < pos-1 {
//			l = mid
//		} else {
//			break
//		}
//	}
//	// l +1 == r
//
//	if p1 == len(nums1) {
//		return nums2[p2]
//	}
//	if p2 == len(nums2) {
//		return nums1[p1]
//	}
//
//	if nums1[p1] < nums2[p2] {
//		return nums1[p1]
//	}
//	return nums2[p2]
//}
//
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	m := len(nums1)
//	n := len(nums2)
//
//	if (m+n)&1 == 1 {
//		p := (m + n) >> 1
//		return float64(searchAposition(nums1, nums2, p+1))
//	}
//
//	p1 := (n + m) >> 1
//	p2 := p1 + 1
//	return float64(searchAposition(nums1, nums2, p1)+searchAposition(nums1, nums2, p2)) / 2
//}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	}

	midIndex1, midIndex2 := totalLength/2-1, totalLength/2
	return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
}

/***

相当于 get KthElement 的 2D 版本吗？


这道题，你直接用 二分查找，就是一个死啊， 有重复的元素。

 */
func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
