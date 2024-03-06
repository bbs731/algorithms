package 定长滑动窗口

/***
这道题， 用前缀和来解， 更好！
 */

// 春雷前缀和的答案：

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	ps := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ps[i] = ps[i-1] + cardPoints[i-1]
	}

	ans := 0
	for i := 0; i <= k; i++ {
		ans = max(ans, ps[i]-ps[0]+ps[n]-ps[n-k+i])
	}
	return ans
}

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	l := []int{}
	l = append(l, cardPoints...)
	l = append(l, cardPoints...)

	if k >= n {
		sum := 0
		for _, x := range cardPoints {
			sum += x
		}
		return sum
	}
	ans := 0
	left := n - k
	right := n - 1
	for left <= n {
		sum := 0
		for i := left; i <= right; i++ { // 这里可以优化 sum
			sum += l[i]
		}
		ans = max(ans, sum)
		left++
		right++
	}
	return ans
}

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	l := []int{}
	l = append(l, cardPoints...)
	l = append(l, cardPoints...) // 其实没有必要， 空间上，是可以优化的。

	if k >= n {
		sum := 0
		for _, x := range cardPoints {
			sum += x
		}
		return sum
	}
	ans := 0
	left := n - k
	right := n - 1
	sum := 0
	for i := left; i < right; i++ { // 注意，这里用的是 i < right 就是先计算  k-1 个数的和， 为了 for loop 里面写起来方便。
		sum += l[i]
	}
	for left <= n {
		sum += l[right]
		ans = max(ans, sum)
		sum -= l[left]
		left++
		right++
	}
	return ans
}

// 灵神的答案：
// 确实是更加的简洁！

func maxScore(cardPoints []int, k int) int {
	s := 0
	for _, x := range cardPoints[:k] {
		s += x
	}
	ans := s
	for i := 1; i <= k; i++ {
		s += cardPoints[len(cardPoints)-i] - cardPoints[k-i]
		ans = max(ans, s)
	}
	return ans
}
