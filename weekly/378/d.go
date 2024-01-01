package weekly

import "fmt"


/*
思路：
1. 需要判断  [0, n/2-1] [n/2,  n-1] 这两段是可以通过变换组成 回文的。
2. 把 [0, n/2 -1] [n/2, n-1] 这两段，不同的地方标记成 nums[] 数组。
3. 把 query [a, b] [c, d] 区间合并， 然后去查 nums[] 看是否覆盖了所有不同。 就可以返回 true

不能简单的合并， 需要检查，对应的 [a, b] 和 [c, d] 对应的字符串是一致的（就是可以通过重新排列，变成相等的字符串)
 */

 /*
 下面代码有问题： 测试用例：   odaxusaweuasuoeudxwa  过不去。
  */

func canMakePalindromeQueries(s string, queries [][]int) []bool {
	n := len(s)
	//occurs_sum := make([]int, n+1)
	//for i:=1; i<=n; i++ {
	//	occurs_sum[i] = occurs_sum[i-1] + int(s[i-1])
	//}

	letters := [26]int{}
	nums := make([]int, n/2)
	total_diffs := 0
	ans := make([]bool, len(queries))
	for i:=0; i<n/2; i++ {
		if s[i] != s[n-i-1] {
			nums[i]	 = 1
			total_diffs++
		}
		letters[s[i]-'a']++
		letters[s[n-i-1]-'a']++
	}
	for i:=0; i<26; i++ {
		if letters[i]%2 !=0 {
			for j :=0; j<len(queries); j++ {
				ans[j] = false
			}
			return ans
		}
	}

	presum := make([]int, n/2+1)
	for i:=1; i<=n/2; i++ {
		presum[i] = presum[i-1] + nums[i-1]
	}
	//total_diffs = presum[n/2] - presum[0]


	for i, q := range queries {
		// [a, b],  [c, d] 把 c, d map 到 [0, n/2-1] 的区间
		a, b, d, c := q[0], q[1],n-q[2]-1, n-q[3]-1


		if a > c{
			a,b ,c ,d = c, d, a, b
		}
		total := 0
		if b < c {
			// 两个区间
			total = presum[b+1] - presum[a] + presum[d+1] - presum[c]

		}else {
			// 合并成一个区间
			total = presum[max(b,d)+1] - presum[a]
		}
		fmt.Println(a, b, c, d, total, total_diffs)
		if total == total_diffs {
			ans[i] = true
		} else {
			ans[i]= false
		}
	}
	return ans
}
