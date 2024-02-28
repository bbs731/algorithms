package one_day_exercise

/****
把每一次， 都变成一个值， 这个策略是不对的！
下面的做法是错误的。
 */
func minIncrements(n int, cost []int) int {
}

func minIncrements(n int, cost []int) int {
	//ans := 0
	//start := 1
	//for 2*start-2 < n-1 {
	//	start = start << 1
	//	tmp := make([]int, 0, start)
	//	for i := start - 1; i <= start-2+start; i++ {
	//		tmp = append(tmp, cost[i])
	//	}
	//	sort.Ints(tmp)
	//	a, b := tmp[start>>1-1], tmp[start>>1]
	//	diffa, diffb := 0, 0
	//	for i := start - 1; i <= start-2+start; i++ {
	//		diffa += abs(a, cost[i])
	//		diffb += abs(b, cost[i])
	//	}
	//	if diffa <= diffb {
	//		ans += diffa
	//	} else {
	//		ans += diffb
	//	}
	//}
	//return ans
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
