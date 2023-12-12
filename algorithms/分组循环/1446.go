package 分组循环

func maxPower(s string) int {
	i := 0
	n := len(s)
	ans := 1
	start := 0

	i++
	for i < n {
		for i < n && s[i] == s[i-1] {
			i++
		}
		ans = max(ans, i-start)
		start = i
		i++
	}
	return ans
}

/*
leetcode 的官方题解：
https://leetcode.cn/problems/consecutive-characters/solutions/1129777/lian-xu-zi-fu-by-leetcode-solution-lctm/

那个？ 更容易理解？
 */

func maxPower(s string) int {
	ans, cnt := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt++
			if cnt > ans {
				ans = cnt
			}
		} else {
			cnt = 1
		}
	}
	return ans
}
