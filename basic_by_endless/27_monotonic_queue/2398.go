package mono_queue

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	n := len(chargeTimes)
	q := make([]int, 0)
	left := 0
	ans := 0
	sum := 0

	for i := 0; i < n; i++ {
		// 入队
		for len(q) > 0 && chargeTimes[i] >= chargeTimes[q[len(q)-1]] {
			q = q[:len(q)-1]
		}

		q = append(q, i)
		sum += runningCosts[i]

		// 出队
		// 左端点不满足，不断右移动
		for len(q) > 0 && chargeTimes[q[0]]+(i-left+1)*sum > int(budget) {
			if q[0] == left {
				q = q[1:]
			}
			sum -= runningCosts[left]
			left++
		}
		ans = max(ans, i-left+1)
	}
	return ans
}

/*
这道题，充分的说明了一个问题，就是， 同一个问题， 不同的写法的复杂度完全不一样。
不但，逻辑变复杂了， 而且，还有特殊情况需要处理
 */
func maximumRobots_chunlei(chargeTimes []int, runningCosts []int, budget int64) int {
	n := len(chargeTimes)
	q := make([]int, 0)
	start := 0
	sum := 0
	ans := 0

	getChargeTime := func(q []int) int {
		if len(q) == 0 {
			return 0
		}
		return chargeTimes[q[0]]
	}

	for i := 0; i < n; i++ {
		// 入队
		total := max(getChargeTime(q), chargeTimes[i]) + (i-start+1)*(sum+runningCosts[i])
		for total > int(budget) && sum > 0 {
			// remove start element
			sum -= runningCosts[start]
			if len(q) != 0 && start == q[0] {
				q = q[1:]
			}
			start++
			total = max(getChargeTime(q), chargeTimes[i]) + (i-start+1)*(sum+runningCosts[i])
		}

		//if sum == 0 && chargeTimes[i] + runningCosts[i] > int(budget) {   // 这个条件也可以
		if total > int(budget) {  // 这里考虑的一个特殊情况就是 item i 就不应该放到队列里, skip to i+1
			start = i + 1
			continue
		}
		// adding robot i
		for len(q) > 0 && chargeTimes[q[len(q)-1]] <= chargeTimes[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		sum += runningCosts[i]

		// 添加 ans
		ans = max(ans, i-start+1)
	}
	return ans
}
