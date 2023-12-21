package _13

func countPairs(coordinates [][]int, k int) int {
	hp :=make(map[[2]int]int)
	ans := 0
	hp[[2]int{coordinates[0][0], coordinates[0][1]}]= 1
	for i:=1; i<len(coordinates); i++ {
		p := coordinates[i]
		for j:=0; j<=100; j++ {  // 枚举 k 的 值域， 因为值域范围够小。O(n) = 5*10^5 * 10^2
			ans += hp[[2]int{ p[0]^j, p[1]^(k-j)}]
		}
		hp[[2]int{p[0], p[1]}]++
	}
	return ans
}
