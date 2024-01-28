package weekly

// 试填法
func minOrAfterOperations(nums []int, k int) int {
	ans := 0
	mask := 0
	for i:=29; i>=0; i-- {
		mask |= 1 << uint(i)
		cnt := 0
		value := -1
		for _, v := range nums {
			value &= v&mask
			if value == 0 {
				value = -1
			}else {
				cnt++
			}
		}
		if cnt > k {
			mask ^= 1 << uint(i) // mask clear the ith bit
			ans |= 1 << uint(i) // for ans ith bit must be on
		}
	}
	return  ans
}
