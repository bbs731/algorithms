package binary_search

import (
	"runtime/debug"
	"sort"
)

/*****

给你两个字符串 s 和 p ，其中 p 是 s 的一个 子序列 。同时，给你一个元素 互不相同 且下标 从 0 开始 计数的整数数组 removable ，该数组是 s 中下标的一个子集（s 的下标也 从 0 开始 计数）。

请你找出一个整数 k（0 <= k <= removable.length），选出 removable 中的 前 k 个下标，然后从 s 中移除这些下标对应的 k 个字符。整数 k 需满足：在执行完上述步骤后， p 仍然是 s 的一个 子序列 。更正式的解释是，对于每个 0 <= i < k ，先标记出位于 s[removable[i]] 的字符，接着移除所有标记过的字符，然后检查 p 是否仍然是 s 的一个子序列。

返回你可以找出的 最大 k ，满足在移除字符后 p 仍然是 s 的一个子序列。

字符串的一个 子序列 是一个由原字符串生成的新字符串，生成过程中可能会移除原字符串中的一些字符（也可能不移除）但不改变剩余字符之间的相对顺序。



示例 1：

输入：s = "abcacb", p = "ab", removable = [3,1,0]
输出：2
解释：在移除下标 3 和 1 对应的字符后，"abcacb" 变成 "accb" 。
"ab" 是 "accb" 的一个子序列。
如果移除下标 3、1 和 0 对应的字符后，"abcacb" 变成 "ccb" ，那么 "ab" 就不再是 s 的一个子序列。
因此，最大的 k 是 2 。
示例 2：

输入：s = "abcbddddd", p = "abcd", removable = [3,2,1,4,5,6]
输出：1
解释：在移除下标 3 对应的字符后，"abcbddddd" 变成 "abcddddd" 。
"abcd" 是 "abcddddd" 的一个子序列。
示例 3：

输入：s = "abcab", p = "abc", removable = [0,1,2,3,4]
输出：0
解释：如果移除数组 removable 的第一个下标，"abc" 就不再是 s 的一个子序列。


提示：

1 <= p.length <= s.length <= 10^5
0 <= removable.length < s.length
0 <= removable[i] < s.length
p 是 s 的一个 子字符串
s 和 p 都由小写英文字母组成
removable 中的元素 互不相同

 */

// 如何判断，一个字符串，是另一个字符串的子串？有什么好的方法吗？
// 春雷， 你选择了一条最难的路？哈哈！ marginal TLE
// 看看下面灵神的代码，如何判断，一个 pattern 是否是一个字符串的  sub-seq 的！
// 但是这道题， 给了我足够的信息，虽然做法上不简洁，但是够复杂， 相当于，在二分中， 又去构造了一个index 的二分结构，然后加循环。
// 这么复杂的逻辑，就错了一点点。sp = tbs[p[i]][ni] + 1 还很快的找到了问题， 一个 Medium 1913 的分数，还不至于如此的复杂。
// 但是至少证明了你，写复杂逻辑的能力。

// 需要强化的知识，判断一个 p 是否是一个 s 的 sub-seq. 有 O(n) 的解法。

func maximumRemovals(s string, p string, removable []int) int {
	debug.SetGCPercent(-1) // 哈哈， 也灭改进，多少， 就是前进了 100ms 左右。
	n := len(removable)
	//fmt.Println(n)
	// 值域 [1, n]
	l, r := 0, n+1

	for l+1 < r {
		mid := (l + r) >> 1

		rem := make(map[int]bool)
		// [0, mid-1] chars in removable is removed.
		for _, v := range removable[:min(mid, n)] {
			rem[v] = true
		}

		tbs := make(map[byte][]int) // we use a table to save each byte's pos in a sorted list
		for i := range s {
			if rem[i] == false {
				tbs[s[i]] = append(tbs[s[i]], i)
			}
		}
		// now judge whether p is sub-seq of s
		judge := true
		sp := 0
		for i := range p {
			ni := sort.SearchInts(tbs[p[i]], sp)
			if ni == len(tbs[p[i]]) {
				judge = false
				break
			}
			sp = tbs[p[i]][ni] + 1 // take the pos. And next char in p going to appear in s with pos > sp
		}

		// 先 true  后 false
		if judge {
			// p is sub-seq
			l = mid
		} else {
			r = mid
		}
	}
	// l+1 = r
	return l
}

// 需要强化的知识，判断一个 p 是否是一个 s 的 sub-seq. 有 O(n) 的解法。

func maximumRemovals(s string, p string, removable []int) int {
	n := len(removable)
	l, r := 0, n+1

	for l+1 < r {
		mid := (l + r) >> 1

		rem := make([]bool, len(s))
		// [0, mid-1] chars in removable is removed.
		for _, v := range removable[:mid] {
			rem[v] = true
		}
		// now judge whether p is sub-seq of s
		// 这个初始化，感觉是怪怪的， 因为是 先true 后 false 的情况， 应该初始化 true 才是自然的想法。
		judge := false
		j := 0

		// 下面这个判断， sub-seq 的逻辑时间复杂度是 O（n+m) 是线性了。
		for i := range s {
			if rem[i] == false && s[i] == p[j] {
				j++
				if j == len(p) {
					judge = true
					break
				}
			}
		}

		// 先 true  后 false
		if judge {
			// p is sub-seq
			l = mid
		} else {
			r = mid
		}
	}
	// l+1 = r
	return l
}

func maximumRemovals(s string, p string, removable []int) int {
	n := len(removable)

	//sort.Search 的使用技巧·其一
	return sort.Search(n, func(x int) bool {
		x++

		rem := make([]bool, len(s))
		// [0, mid-1] chars in removable is removed.
		for _, v := range removable[:x] {
			rem[v] = true
		}

		judge := false
		j := 0

		// 下面这个判断， sub-seq 的逻辑时间复杂度是 O（n+m) 是线性了。
		for i := range s {
			if rem[i] == false && s[i] == p[j] {
				j++
				if j == len(p) {
					judge = true
					//break
					return false
				}
			}
		}
		// 先 true 后 false , 所以这里应该取反。
		return !judge
	})
}

// 灵神的题解：
func maximumRemovals(s, p string, id []int) int {
	return sort.Search(len(id), func(k int) bool {
		del := make([]bool, len(s))
		for _, i := range id[:k+1] {
			del[i] = true
		}
		j := 0
		for i := range s {
			if !del[i] && s[i] == p[j] {
				if j++; j == len(p) {
					return false
				}
			}
		}
		return true
	})
}
