package dp

// 这是一个非彼那切数列
func climbStairs(n int) int {
	l := make([]int, n+1)
	l[0] = 1
	l[1] = 1

	// l[i+2] = l[i+1] + l[i]
	for i := 2; i <= n; i++ {
		l[i] = l[i-1] + l[i-2]
	}
	return l[n]
}

func climbStairs_2(n int) int {
	f0 := 1
	f1 := 1

	// l[i+2] = l[i+1] + l[i]
	for i := 2; i <= n; i++ {
		new_f := f0 + f1
		f0 = f1
		f1 = new_f
	}
	return f1
}

func climbStairs_1(n int) int {
	l := make([]int, n+1)
	l[0] = 1
	l[1] = 1

	// l[i+2] = l[i+1] + l[i]
	for i := 0; i < n-1; i++ {
		l[i+2] = l[i] + l[i+1]
	}
	return l[n]
}
