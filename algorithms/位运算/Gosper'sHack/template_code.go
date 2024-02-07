package gosper_hack

import "math/bits"

func _() {

	loopSubsetK := func(a []int, k int) {
		n := len(a)
		for sub := 1<<k - 1; sub < 1<<n; {
			lb := sub & (-sub)
			x := sub + lb
			// 下式等价于 sub = (sub^x)/lb>>2 | x
			// 把除法改成右移 bits.TrailingZeros 可以快好几倍
			sub = (sub^x)>>bits.TrailingZeros(uint(lb))>>2 | x
		}
	}
}
