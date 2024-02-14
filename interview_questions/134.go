package interview_questions

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	diff := make([]int, 2*n-1)
	for i := range diff {
		diff[i] = gas[i%n] - cost[i%n]
	}
	// now we use sliding window to solve the problem
	left := 0
	sum := 0
	for i := 0; left < n && i < 2*n-1; i++ { // loop right edge of the window
		sum += diff[i]
		for sum < 0 && left <= i {
			sum -= diff[left]
			left++
		}

		if sum >= 0 && i-left+1 == n {
			return left
		}
	}
	return -1
}
