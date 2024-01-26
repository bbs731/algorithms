package loop


// 按照分组循环的模版来写
func getDescentPeriods(prices []int) int64 {
	n := len(prices)
	ans := n

	for i:=0; i<n; {
		start := i

		// 看下面的 for loop  i 先 ++ 之后再处理也是可以的。
		for i++; i<n && prices[i]+1 == prices[i-1]; i++ {
			//这里有陷阱啊！
			//在 for loop 里面 break, i++ 是不执行的!
		}
		l := i-1 - start + 1
		ans += l*(l-1)/2
	}
	return int64(ans)
}
