package binaryIndexTree

type BIT struct {
	n int
	C []int // index 从1 开始，  1，2，... n  // 初始化全为0, 所以可以理解为 a[1], a[2], .... a[n] 全为0，然后利用 add()  设置为原数组的值， 譬如 add(1, a[1]), add(2, a[2])
}

func lowbit(x int) int {
	return x & (-x)
}
func (b *BIT) add(x int, k int) { // a[x] + k
	for x <= b.n {
		b.C[x] = b.C[x] + k
		x += lowbit(x)
	}
}

func (b *BIT) getsum(x int) int { // 求前缀和 a[1], a[2], ... a[x]
	ans := 0
	for x >= 1 {
		ans = ans + b.C[x]
		x -= lowbit(x)
	}
	return ans
}
