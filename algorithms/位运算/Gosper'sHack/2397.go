package gosper_hack

import "math/bits"

func maximumRows(matrix [][]int, numSelect int) int {
	m := len(matrix)
	n := len(matrix[0])

	mask := make([]int, m)
	for i, r := range matrix {
		for j, c := range r {
			mask [i] |= c << uint(j)
		}
	}

	ans := 0
	// Gosper's 枚举子集的算法 一共 C(n, k) 种
	for sub := 1<<numSelect - 1; sub < 1<<n; {
		cnts := 0
		for i := range mask {
			if mask[i]&sub == mask[i] {
				cnts++
			}
		}
		ans = max(ans, cnts)
		lowbit := sub & (-sub)
		x := sub + lowbit
		sub = (sub^x)>>bits.TrailingZeros(uint(lowbit))>>2 | x
	}
	return ans
}
