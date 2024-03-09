package montonic_stack

/***
这题目，竟然可以用单调栈来解决。

这个思路，太难想了， 一看到连续子数组， 就想到模拟，想过 ST table 想过 BIT。  没想过 单调栈。
 */

func sumSubarrayMins(arr []int) int {
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		left[i] = -1
		right[i] = n
	}
	st := []int{}
	for i := 0; i < n; i++ {
		a := arr[i]
		for len(st) > 0 && a < arr[st[len(st)-1]] {
			// pop stack
			top := st[len(st)-1]
			st = st[:len(st)-1]
			right[top] = i
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	MOD := int(1e9) + 7
	ans := 0
	for i := 0; i < n; i++ {
		ans += (i - left[i]) * (right[i] - i) * arr[i]
		ans = ans % MOD
	}

	return ans % MOD
}
