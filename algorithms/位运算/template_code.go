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

	//遍历集合： 假设元素包括 从 0 到 n-1 个元素
	for i := 0; i < n; i++ {
		if s>>i&1 == 1 { // i 在 s 中
			// 处理 i 的逻辑
		}
	}

	//枚举集合
	for s := 0; s < 1<<n; s++ {
		// 处理s 的逻辑, s 是一个集合里面的元素范围是 0 到 n-1
	}

	//从大到小枚举集合 s 的所有非空子集 sub:
	for sub := s; sub > 0; sub = (sub - 1) & s { // sub = (sub-1)&s 这个有深意，看上面灵神链接里面的解释
		// 处理 子集 sub
	}
	//从大到小枚举集合 s 的所有子集sub 包括空集的情况：
	for sub := s; ; {
		// 处理子集 sub
		sub = (sub - 1) & s
		if sub == s {
			break
		}
	}

}
