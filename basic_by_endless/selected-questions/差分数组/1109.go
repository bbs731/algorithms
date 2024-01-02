package array

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n)

	for _, book := range bookings {
		s, e , seats := book[0], book[1], book[2]
		diff[s-1] +=seats
		if e < n {
			diff[e] -= seats
		}
	}

	//presum := make([]int, n)  // 优化空间复杂度，可以不创建 presum 数组。
	//presum[0] = diff[0]
	for i:=1;i<n;i++ {
		//presum[i] = presum[i-1] + diff[i]
		diff[i] += diff[i-1]
	}
	return diff
}
