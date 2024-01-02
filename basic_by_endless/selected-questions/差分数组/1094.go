package array

func carPooling(trips [][]int, capacity int) bool {
	N := 1000
	nums := make([]int, N+2)

	for _, trip := range trips {
		d, s, e := trip[0], trip[1], trip[2]
		nums[s] += d
		nums[e] -=d
	}

	presum := make([]int, N+1)
	presum[0] = nums[0]
	for i:=1; i<=N; i++ {
		presum[i]= presum[i-1] + nums[i]
		if presum[i] > capacity {
			return false
		}
	}

	//for i:=0; i<=N;i++ {
	//	if presum[i] > capacity {
	//		return false
	//	}
	//}
	return true

}
