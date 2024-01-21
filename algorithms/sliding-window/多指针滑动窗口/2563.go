package sliding_window

import (
	"sort"
)

/***

[0, 1, 4, 4, 5, 7]

 */

 /***
 这道题，做了3个小时吗？ 怎么办？

 lower_bound,
 upper_bound 还是心魔啊！
  */


// 春雷你对 upper_bound 的理解，一直是有错误的。

func upper_bound(nums []int, v int, r int) int {
	left :=  -1
	right := r

	for left +1 < right {
		mid := (left + right)/2
		if nums[mid] <= v {
			left = mid
		} else {
			right = mid
		}
	}
	// is this upper bound?
	//return left
	return right
}

func lower_bound(nums []int, v int,  r int) int {
	left := -1
	right := r

	for left + 1 < right {
		 mid := (left+right)/2
		 if nums[mid] < v {
		 	left = mid
		 } else {
		 	right = mid
		 }
	}
	return right
}


func countFairPairs(nums []int, lower int, upper int) int64 {
	n := len(nums)
	sort.Ints(nums)


	ans := 0
	for right:=0; right< n; right++  {
		v := nums[right]

		pos_right:= upper_bound(nums, upper-v, right)
		//pos_right:= sort.SearchInts(nums[:right], upper-v+1)
		pos_left := lower_bound(nums, lower -v, right)
		ans += pos_right - pos_left
	}
	return int64(ans)
}

func countFairPairs(nums []int, lower int, upper int) int64 {
	n := len(nums)
	sort.Ints(nums)


	ans := 0
	for left :=0; left< n-1; left++  {
		v := nums[left]

		pos_right:= upper_bound(nums, upper-v, left+1)
		pos_left := lower_bound(nums, lower -v, left+1  )

		if pos_right >= pos_left {
			ans += pos_right - pos_left + 1
		}
	}
	return int64(ans)
}

func countFairPairs(nums []int, lower int, upper int) int64 {
	n := len(nums)
	sort.Ints(nums)

	ans := 0
	for left :=0; left< n-1; left++  {
		v := nums[left]

		for k:=n-1; k>left; k-- {
			if nums[k] +v <=upper && nums[k]+v >= lower {
				ans++
			}
		}
	}

	return int64(ans)
}

//
////  这都写的他妈的什么鬼！
//func countFairPairs(nums []int, lower int, upper int) int64 {
//	n := len(nums)
//	sort.Ints(nums)
//
//	ans := 0
//	for left :=0; left< n-1; left++  {
//		v := nums[left]
//
//		if v + nums[n-1] < lower {
//			continue
//		}
//
//		if v + nums[n-1] > upper {
//			break
//		}
//
//		if v + nums[left+1] > lower {
//			break
//		}
//
//
//		for k:=n-1; k>left; k-- {
//			if nums[k] +v <=upper && nums[k]+v >= lower {
//				ans++
//			}
//		}
//	}
//
//	return int64(ans)
//}
