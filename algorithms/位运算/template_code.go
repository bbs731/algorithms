package bits_operation

import "math/bits"

// 总结在这里:  https://leetcode.cn/circle/discuss/CaOJ45/

func _() {
	var n int
	var s uint
	// 核心的思想就是，把数字 对应成一个集合 s, 对 bit 的操作，就是对集合的操作。
	// 下面介绍一些，基础的操作，在一个 uint （64bit) 个数之内的， 如果超过了看进阶的  type Bitset []uint 的操作
	// https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/bits.go#L307
	// Set 元素的 index 是从 0 开始的, 最大元素是 bits.Len(s) - 1。

	//集合的大小
	bits.OnesCount(s)

	//集合中的最大元素
	_ = bits.Len(s) - 1

	//集合中的最小元素
	_ = bits.TrailingZeros(s)

	//二进制最低 1 及其后面的 0 叫做 lowbit
	lowbit := s & (-s) // lowbit == 1<<bits.TrailingZeros(s)  除了 s=0 的情况

	//遍历集合：
	for i := 0; i < n; i++ {
		if s>>i&1 == 1 { // i 在 s 中
			// 处理 i 的逻辑
		}
	}

}
