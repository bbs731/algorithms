package binary_search

// 用红蓝染色法， 来寻找答案。
func mySqrt(x int) int {
	l, r := 0, x

	for l <= r {
		mid := l + (r - l) / 2
		if mid * mid <= x {
			l = mid + 1   //  循环不变量  <=l-1 都是红色
		} else {
			r = mid - 1   //  >= r+1 都是蓝色
		}
	}
	return r // l = r + 1

}

// 上面二分的方法，比这个 O(n) 的好。 
func mySqrt(x int) int {

	for i:=0; true ; i++{
		n := i*i
		if n == x{
			return i
		}
		if n > x {
			return i-1
		}
	}
	return 0
}
