package one_day_exercise

func change(amount int, coins []int) int {
	f := make([]int, amount+1)
	f[0] = 1

	// loop coins 和 amount 的顺序，有决定性的意义！
	// 为什么最外层的 loop, 需要先loop coins？
	for _, c := range coins {
		for i := 1; i <= amount; i++ {
			if i >= c {
				f[i] += f[i-c]
			}
		}
	}
	//fmt.Println(f)
	return f[amount]
}
