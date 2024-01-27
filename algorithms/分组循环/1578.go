package loop

/***
Horry Template 分组循环
 */
func minCost(colors string, neededTime []int) int {
	n := len(colors)
	ans := 0

	for i:=0; i<n; {
		start := i
		mv := 0
		sum := 0
		for ; i<n && colors[i] == colors[start]; i++ {
			sum += neededTime[i]
			mv = max(mv, neededTime[i])
		}
		ans += sum - mv
	}
	return ans
}
