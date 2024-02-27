package binaryIndexTree

/*

灵神的题解：

https://leetcode.cn/problems/range-sum-query-mutable/solutions/2524481/dai-ni-fa-ming-shu-zhuang-shu-zu-fu-shu-lyfll/
 */

type NumArray struct {
	original []int
	b        []int
	n        int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	b := make([]int, n+1)

	// construct BIT in O(n) time  // 这里不熟悉， 容易出问题。
	for i := 0; i < n; i++ {
		b[i+1] += nums[i] // 这里有bug 需要 +=,   = 是不对的，需要累加
		c := i + 1 + lowbit(i+1)
		if c <= n {
			b[c] += b[i+1] // update its parent, 需要 +=  child b[i+1] 而不是 num[i]
		}
	}

	return NumArray{
		nums,
		b,
		n,
	}
}

func lowbit(x int) int {
	return x & (-x)
}

func (this *NumArray) Update(index int, val int) {
	v := val - this.original[index]

	x := index + 1
	for x <= this.n {
		this.b[x] += v
		x += lowbit(x)
	}

	// update original string
	this.original[index] = val
}

func (this *NumArray) query(x int) int {
	ans := 0
	for x >= 1 {
		ans += this.b[x]
		x -= lowbit(x)
	}
	return ans
}
func (this *NumArray) SumRange(left int, right int) int {
	return this.query(right+1) - this.query(left)
}
