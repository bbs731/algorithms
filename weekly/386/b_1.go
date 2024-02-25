package weekly


func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	n := len(bottomLeft)

	ans := 0
	for i:=0; i<n; i++ {
		for j:=i+1;j<n; j++ {
			// 求出新的 ， left-bottom 和 top-right 的节点。
			bl1, bl2 := bottomLeft[i], bottomLeft[j]
			blx, bly := max(bl1[0], bl2[0]), max(bl1[1], bl2[1])

			tr1,tr2 := topRight[i], topRight[j]
			trx, try:= min(tr1[0], tr2[0]), min(tr1[1], tr2[1])

			ans = max(ans, min(max(trx-blx, 0), max(try-bly, 0)))
		}
	}
	return int64(ans*ans)
}



func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
