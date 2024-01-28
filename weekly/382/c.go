package weekly


// [1, n] 之间有多少个偶数
func numsofEvens(n int) int {
	if n <= 1 {
		return 0
	}
	return n/2
}

func flowerGame(n int, m int) int64 {
	ans := 0
	for i := 1; i<=n; i++ {
		//if i & 1 == 0 {
		//	// even
		//	ans += m - m/2
		//
		//} else {
		//	// odd
		//	ans += m/2
		//}
		ans += m/2
	}

	// return n*m/2
	return int64(ans)
}
