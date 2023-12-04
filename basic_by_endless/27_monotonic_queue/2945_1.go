package mono_queue

/*
1,2,1,3,3
1,2,7  贪心得到错误的答案
1,3,3,3 这个是正确的答案。


f[i] 定义为， 到 ith element 的时候， 能形成的最长升序的长度。
那么考虑， f[i] 是如何通过其它状态（之前的状态）转移过来的， j < i

f[i] = f[j] + 1  and  sum[j+1 .. i] >= last[j]    s[i]- s[j] >= last[j]  -> s[i] >= s[j] +last[j]
其中 last 数组表示， last[i] 到 i 位置，形成的升序子串的最后一个element 的值。

f[i] = max(f[j]) + 1

1. 决定性的式子是 s[i] >= s[j] + last[j]
2. 并且， f[i] 是递增的。
根据这两个性质，写出的单调队列的逻辑。

 */

func findMaximumLength(nums []int) int {
	n := len(nums)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + nums[i-1]
	}
	q := []int{0}
	f := make([]int, n+1)
	last := make([]int, n+1)

	for i := 1; i <= n; i++ {
		for len(q) > 1 && s[q[1]]+last[q[1]] <= s[i] {
			q = q[1:]
		}

		f[i] = f[q[0]] + 1
		last[i] = s[i] - s[q[0]]

		for len(q) > 0 && s[q[len(q)-1]]+last[q[len(q)-1]] >= s[i]+last[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	return f[n]
}
