package binaryIndexTree

type BIT struct {
	n int
	a []int

	// index 1 to n
	C []int // index 从1开始，原始数组的index 从0开始 a[0], a[1],.... a[n-1] 所以需要考虑下标不统一的问题，具体看建表！
}

func lowbit(x int) int {
	return x & (-x)
}

// 这里的 x 下标指代的是 b.C 的从1开始的下标， 不是指代的原始数组 a 的下标。
// 对应的操作，应该是给  a[x-1] 这个数 +k
func (b *BIT) add(x int, k int) {
	for x <= b.n {
		b.C[x] = b.C[x] + k
		x += lowbit(x)
	}
}

// 相当于再求， a[0],......a[x-1] 的前缀和。 调用的时候使用 b.getsum(x+1)
// 这里的 x 在 BIT的 C 数组里看，是下标从 1 开始的。 需要考虑和原始数组 a[] 之间的下标变换关系。
func (b *BIT) getsum(x int) int { //这里的 x 是按照 C 的下标来考虑的, 就是 index 从 1 开始考虑的。
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
func (b *BIT) init(a []int) {
	b.C = make([]int, b.n+1)
	for i := 1; i <= b.n; i++ {
		b.C[i] += a[i-1] // 这里 a[] 数组下标 从 0 开始，所以需要使用 a[i-1]
		j := i + lowbit(i)
		if j <= b.n {
			b.C[j] += b.C[i]
		}
	}
}

// 第二种方法，根据定义
func (b *BIT) int(sum []int) {
	for i := 1; i <= b.n; i++ {
		//根据定义
		b.C[i] = sum[i] - sum[i-lowbit(i)]
	}
}

// 求区间最大值的 BIT 做法， 时间复杂度是 logn*logn
// 证明在这里：  https://oi-wiki.org/ds/fenwick/#%E5%8C%BA%E9%97%B4%E6%9F%A5%E8%AF%A2_1
func (b *BIT) getMax(l, r int) int {
	ans := 0
	for r >= l {
		ans = max(ans, b.a[r-1])
		r--
		for ; r-lowbit(r) >= l; r -= lowbit(r) {
			ans = max(ans, b.C[r])
		}
	}
	return ans
}

func (b *BIT) init(a []int) {
	n := len(a)
	b.a, b.n = a, n
	b.C = make([]int, n+1)

	for i := 1; i <= n; i++ {
		b.C[i] = max(b.C[i], a[i-1])
		if j := i + lowbit(i); j <= n {
			b.C[j] = max(b.C[j], b.C[i]) // 这里经常出 bug. 应该用 b.C[i] 来更新它的parent, 不要用 a[i-1]
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
