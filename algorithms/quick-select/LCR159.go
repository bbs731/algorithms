package quick_select

import "math/rand"

/***
这道题， 充分，的说明了， cnt 的 index 是需要从 0 开始的。
如果传进来的 cnt index 是从 1 开始的，需要，特殊的处理一下。

 */
func inventoryManagement(stock []int, cnt int) []int {
	n := len(stock)
	//cnt--

	l, r := 0, n-1
	rand.Shuffle(n, func(i, j int) {
		stock[i], stock[j] = stock[j], stock[i]
	})
	for l < r {
		i, j := l, r+1
		v := stock[l]

		for {
			for i++; i < r && stock[i] < v; i++ {
			}
			for j--; j > l && stock[j] > v; j-- {
			}
			if i >= j {
				break
			}
			stock[i], stock[j] = stock[j], stock[i]
		}
		stock[l], stock[j] = stock[j], stock[l]
		if j == cnt-1 {
			return stock[:cnt]
		} else if j > cnt-1 {
			r = j - 1
		} else {
			l = j + 1
		}
	}
	return stock[:cnt]
}
