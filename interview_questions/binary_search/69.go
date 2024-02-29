package binary_search

func mySqrt(x int) int {
	l, r := -1, x+1
	for l+1 < r {
		mid := l + (r-l)>>1
		if mid*mid > x {
			r = mid
		} else {
			l = mid
		}
	}
	return l
}
