package binaryIndexTree

import "sort"

/***
和 LCR 170 是一类的题目。 不是求总数， 是求单独 postion 的个数。 就是没个 query 的结果，都需要记录
 */

type BIT struct {
	b []int
	n int
}

func lowbit(x int) int {
	return x & (-x)
}

func (b *BIT) add(x, v int) {
	for x <= b.n {
		b.b[x] += v
		x += lowbit(x)
	}
}

func (b *BIT) query(x int) int {
	ans := 0
	for x >= 1 {
		ans += b.b[x]
		x -= lowbit(x)
	}
	return ans
}

func countSmaller(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	//tmp := []int{} // 这是一个bug, 如果，tmp 初始化成这个样子，那么， 不会 copy nums to tmp, 因为 tmp 大小为0
	tmp := make([]int, n)
	copy(tmp, nums)
	sort.Ints(tmp)

	// 这个相当于，在做 nums 数组值域的离散化. 这里比 LCR170 处理的简单且有技巧！但是这个不是通用的逻辑。
	for i, r := range nums {
		nums[i] = sort.SearchInts(tmp, r)
	}
	b := BIT{
		make([]int, n+1),
		n,
	}

	for i := n - 1; i >= 0; i-- {
		ans[i] = b.query(nums[i])
		b.add(nums[i]+1, 1)
	}
	return ans
}
