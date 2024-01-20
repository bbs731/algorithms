package sliding_window


// 这要是作为面试题， 死 100次了。
// 这个题，不可能 1300 多分。
func countCompleteSubarrays(nums []int) int {
	//n :=len(nums)
	cnts := make(map[int]struct{})
	for _, x :=range nums{
		cnts[x] = struct{}{}
	}
	m := len(cnts)


	ans := 0
	left := 0
	c := make(map[int]int)
	// 我操， 这个太难！
	for _, v :=range nums {
		c[v]++
		for len(c) == m {
			x := nums[left]
			c[x]--
			if c[x]==0 {
				delete(c, x)
			}
			left++
		}
		ans += left
	}
	return ans
}

// 这题要是想不到方法，就太难！
func countCompleteSubarrays(nums []int) int {
	n := len(nums)
	cnts := make(map[int]int, n)

	for _, x := range nums {
		cnts[x]++
	}

	if len(cnts) == 1 {
		return n * (n + 1) / 2
	}

	//total := len(cnts)
	left := 0
	right := n - 1

	ans := 1
	for left <= right {
		l := nums[left]
		r := nums[right]
		if cnts[l] > 1 && cnts[r] > 1 {
			if l == r && cnts[l] == 2 {
				ans += 2
			} else {
				ans += 3
			}
			cnts[l]--
			cnts[r]--
			left++
			right--
		} else if cnts[l] > 1 {
				ans +=1
			cnts[l]--
			left++
		} else if cnts[r] > 1 {
				ans +=1
			cnts[r]--
			right--

		} else {
			break
		}
	}
	return ans
}
