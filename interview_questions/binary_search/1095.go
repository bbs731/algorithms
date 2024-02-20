package binary_search

import "sort"

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

 /***
 借助 162 的题解，先找到峰值。

赞啊， 没想到有 162 打底的话，是不是就废了？
  */
func findInMountainArray(target int, mountainArr *MountainArray) int {
	n := mountainArr.length()
	l , r := 0, n-1
	for l + 1 < r {
		mid := (l+r)>>1
		if mountainArr.get(mid) < mountainArr.get(mid+1) {
			l = mid
		} else {
			r = mid
		}
	}
	// l +1 == r
	// r is now the peak
	peak := r
	if mountainArr.get(peak) == target {
		return peak
	}

	l , r = -1, peak
	for l + 1 < r {
		mid := (l+r)>>1
		mv := mountainArr.get(mid)
		if mv == target {
			return mid
		} else if  mv > target {
			r = mid
		} else {
			l = mid
		}
	}

	// search right part
	l , r = peak, n
	for l +1 < r {
		mid := (l+r)>>1
		mv := mountainArr.get(mid)
		if mv == target {
			return mid
		} else if mv > target {
			l = mid
		} else {
			r = mid
		}
	}
	return -1

}
