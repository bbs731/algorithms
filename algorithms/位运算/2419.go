package bits_operation

/***
这道题，好难啊！ WA 了至少三次了。
 */

func longestSubarray(nums []int) int {
	max_num := 0
	ans := 1
	cnts := 0

	for _, num := range nums {
		if num > max_num {
			// need to reset
			max_num = num
			cnts = 1
			ans = 1
		} else if num == max_num {
			cnts++
			ans = max(ans, cnts)
		} else {
			cnts = 0
		}
	}
	return ans // 连续最大数的个数。
}

/***
我kao 好难的写法， 错了能有5次 WA
 */
func longestSubarray(nums []int) int {
	max_num := 0
	ans := 1
	cnts := 0

	for i, num := range nums {
		if num > max_num {
			// need to reset
			max_num = num
			cnts = 1
			ans = 1
		} else if num == max_num {
			if nums[i-1] == max_num {
				cnts++
				ans = max(ans, cnts)
			} else {
				cnts = 1
			}
		}
	}
	return ans // 连续最大数的个数。
}
