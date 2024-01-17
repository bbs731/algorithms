package binaryIndexTree

type BIT struct {
	n int
	C []int // index 从1 开始，  1，2，... n  // 初始化全为0, 所以可以理解为 a[1], a[2], .... a[n] 全为0，然后利用 add()  设置为原数组的值， 譬如 add(1, a[1]), add(2, a[2])
	// index 1 to n
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

// O(n) 建树  https://oi-wiki.org//ds/fenwick/
//两种方法。

// 第一种方法，还是有点技巧的。 （更新自身节点和父亲节点）
func (b *BIT) init( a[]int) {
	b.C = make([]int, b.n+1)
	for i:=1; i<=b.n; i++ {
		b.C[i] += a[i]
		j := i+lowbit(i)
		if j <= b.n {
			b.C[j] += b.C[i]
		}
	}
}
// 第二种方法，根据定义
func (b *BIT) int(sum []int) {
	for i:=1; i<=b.n; i++{
		//根据定义
		b.C[i] = sum[i]-sum[i-lowbit(i)]
	}
}