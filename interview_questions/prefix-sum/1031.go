package prefix_sum

/****
这是最朴素的算法， 过不了面试。
 */
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	n := len(nums)

	// left firstLen <= secondLen
	if firstLen > secondLen {
		firstLen, secondLen = secondLen, firstLen
	}

	first := make([]int, n-firstLen+1)
	second := make([]int, n-secondLen+1)
	ans := 0

	for i := 0; i < firstLen; i++ {
		first[0] += nums[i]
	}
	for i := 1; i <= n-firstLen; i++ { // <= n-firstLen  还是  < n-firstLen 这里太容易错了。
		first[i] = first[i-1] - nums[i-1] + nums[i+firstLen-1]
	}
	for i := 0; i < secondLen; i++ {
		second[0] += nums[i]
		ans = second[0]
	}
	for i := 0; i <= n-secondLen; i++ {
		if i > 0 {
			second[i] = second[i-1] - nums[i-1] + nums[i+secondLen-1]
		}

		for j := 0; j <= i-firstLen; j++ { // 不相交，不相交！ 计算端点
			ans = max(ans, first[j]+second[i])
		}

		for j := i + secondLen; j <= n-firstLen; j++ { // 不相交。 计算端点
			ans = max(ans, first[j]+second[i])
		}
	}
	return ans
}

/***
虽然用  BIT 做了优化。 但是 BIT  的 getMax 是 logn*logn 的，还是过不了面试，
还有 O(n)的解法。 去看题解。
 */
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	n := len(nums)

	// left firstLen <= secondLen
	if firstLen > secondLen {
		firstLen, secondLen = secondLen, firstLen
	}

	first := make([]int, n-firstLen+1)
	second := make([]int, n-secondLen+1)
	ans := 0

	for i := 0; i < firstLen; i++ {
		first[0] += nums[i]
	}
	for i := 1; i <= n-firstLen; i++ { // <= n-firstLen  还是  < n-firstLen 这里太容易错了。
		first[i] = first[i-1] - nums[i-1] + nums[i+firstLen-1]
	}
	for i := 0; i < secondLen; i++ {
		second[0] += nums[i]
		ans = second[0]
	}

	b := &BIT{}
	b.init(first)

	for i := 0; i <= n-secondLen; i++ {
		if i > 0 {
			second[i] = second[i-1] - nums[i-1] + nums[i+secondLen-1]
		}

		//for j := 0; j <= i-firstLen; j++ { // 不相交，不相交！ 计算端点
		//	ans = max(ans, first[j]+second[i])
		//}
		ans = max(ans, second[i]+b.getMax(1, i-firstLen+1))
		//for j := i + secondLen; j <= n-firstLen; j++ { // 不相交。 计算端点
		//	ans = max(ans, first[j]+second[i])
		//}
		ans = max(ans, second[i]+b.getMax(i+secondLen+1, n-firstLen+1))
	}
	return ans
}

/***
可以用 BIT来优化  first 数组， 来支持 range query,  我们先来试一下用 BIT来优化，复杂度是 logn 的，

之后我们在用 ST table 来优化一下， 是 O(1)的 query 时间，但是建表的时间是 O(n*lgn) 的， 所以，其实相对  BIT 也没有太多的优势。
 */

type BIT struct {
	a []int
	C []int
	n int
}

func lowbit(x int) int {
	return x & (-x)
}

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

/****

https://leetcode.cn/problems/maximum-sum-of-two-non-overlapping-subarrays/solutions/2245647/tu-jie-mei-you-si-lu-yi-zhang-tu-miao-do-3lli/

这是灵神给的答案， 太他们的巧妙了！ 这个就叫做尺取法吧!
 */

func maxSumTwoNoOverlap(nums []int, firstLen, secondLen int) (ans int) {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x // 计算 nums 的前缀和
	}
	f := func(firstLen, secondLen int) {
		maxSumA := 0
		for i := firstLen + secondLen; i <= n; i++ {
			maxSumA = max(maxSumA, s[i-secondLen]-s[i-secondLen-firstLen])
			ans = max(ans, maxSumA+s[i]-s[i-secondLen])
		}
	}
	f(firstLen, secondLen) // 左 a 右 b
	f(secondLen, firstLen) // 左 b 右 a
	return
}
