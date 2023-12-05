package dp

/*
这个代码，看起来更清爽吧， 冲 300 LIS 改过来， 增加 t 数组，用来统计数量。
 */
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	t := make([]int, n)
	ans := 1
	total := 0

	for i, x := range nums {
		f[i] = 1
		t[i] = 1
		for j := 0; j < i; j++ {
			if x > nums[j] {
				if f[j]+1 > f[i] {
					t[i] = t[j]
				} else if f[j]+1 == f[i] {
					t[i] += t[j]
				}
				f[i] = max(f[i], f[j]+1)
			}
		}
		ans = max(ans, f[i])
	}

	for i := range nums {
		if f[i] == ans {
			total += t[i]
		}
	}
	return total
}

// f[i] = max(f[j]) + 1
//[1,2,4,3,5,4,7,2]
// [1,2,3,1,2,3,1,2,3] 这个测试用例绝绝子。
// 这道题，充分的说明了你的脑子不是很好使
func findNumberOfLIS(nums []int) int {
	ans := 0
	longest := 1
	f := make([]int, len(nums))
	t := make([]int, len(nums))

	for i, x := range nums {
		f[i] = 1
		t[i] = 0
		local := 0
		for j := 0; j < i; j++ {
			if x > nums[j] {
				f[i] = max(f[i], f[j]+1)
				if f[j]+1 > local {
					local = f[i]
					t[i] = t[j]
				} else if f[j]+1 == local {
					t[i] += t[j]
				}
			}
			longest = max(longest, f[i])
		}
		if t[i] == 0 {
			t[i] = 1
		}
	}

	for i := range nums {
		if f[i] == longest {
			ans += t[i]
		}
	}
	return ans
}
