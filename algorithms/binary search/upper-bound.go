package main

import (
	"fmt"
	"sort"
)

/*
去看一下 upper_bound 的定义：
https://stackoverflow.com/questions/28389065/difference-between-basic-binary-search-for-upper-bound-and-lower-bound
 */

func main() {

	num := 5
	nums := []int{1, 1, 2, 4, 5, 5, 5, 7, 7, 8, 9}

	var Search func(int, func(int) bool) int
	// upper-bound search
	Search = func(n int, f func(int) bool) int {
		// Define f(-1) == false and f(n) == true.
		// Invariant: f(i-1) == false, f(j) == true.
		i, j := 0, n
		for i < j {
			h := int(uint(i+j) >> 1) // avoid overflow when computing h
			// i ≤ h < j
			if !f(h) {
				i = h + 1 // preserves f(i-1) == false
			} else {
				j = h // preserves f(j) == true
			}
		}

		return i
	}

	lower_bound := sort.Search(len(nums), func(h int) bool {
		return nums[h] >= num
	})
	fmt.Println(lower_bound)

	upper_bound := sort.Search(len(nums), func(h int) bool {
		return nums[h] > num
	})
	fmt.Println(upper_bound)

	upper_bound = Search(len(nums), func(h int) bool {
		return nums[h] > num
	})
	fmt.Println(upper_bound)

}
