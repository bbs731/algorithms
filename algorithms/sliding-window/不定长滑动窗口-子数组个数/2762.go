package sliding_window

//灵神的题解：
//https://leetcode.cn/problems/continuous-subarrays/solutions/2327219/shuang-zhi-zhen-ping-heng-shu-ha-xi-biao-4frl/
//

// 这次我觉得，我的题解也不错！ 虽然varibale 的命名规则有待商榷。
func continuousSubarrays(nums []int) int64 {
	left := 0
	n := len(nums)
	ans := 0

	m := make(map[int]int) // num and its last appeared pos

	for i:=0; i<n; i++ {
		v := nums[i]
		m[v] = i

		newLeft := -1
		toDelete := []int{}
		for k, pos := range m {
			if v < k-2 || v > k+2 {
				newLeft = max(newLeft, pos)  // 保存left需要更新的值。
				toDelete = append(toDelete, k)
			}
		}
		if newLeft != -1 {
			left = newLeft+1  // 这里容易错误！
			for _, k := range toDelete {
				delete(m, k)
			}
		}
		ans += i-left + 1
	}
	return int64(ans)
}



// 会超时， 如何优化？
func continuousSubarrays(nums []int) int64 {
	left := 0
	n := len(nums)
	ans := 1

	for i:=1; i<n; i++ {
		v := nums[i]
		j := i-1;
		for ; j>=left; j-- {
			if v <nums[j]-2  || v > nums[j] +2 {
				break
			}
		}
		j++
		if j != left {
			left = j
		}
		ans += i-left + 1
	}
	return int64(ans)
}
