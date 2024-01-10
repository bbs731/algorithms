package _600_1900

/*

给定由 n 个字符串组成的数组 strs，其中每个字符串长度相等。

选取一个删除索引序列，对于 strs 中的每个字符串，删除对应每个索引处的字符。

比如，有 strs = ["abcdef", "uvwxyz"]，删除索引序列 {0, 2, 3}，删除后 strs 为["bef", "vyz"]。

假设，我们选择了一组删除索引 answer，那么在执行删除操作之后，最终得到的数组的元素是按 字典序（strs[0] <= strs[1] <= strs[2] ... <= strs[n - 1]）排列的，然后请你返回 answer.length 的最小可能值。



示例 1：

输入：strs = ["ca","bb","ac"]
输出：1
解释：
删除第一列后，strs = ["a", "b", "c"]。
现在 strs 中元素是按字典排列的 (即，strs[0] <= strs[1] <= strs[2])。
我们至少需要进行 1 次删除，因为最初 strs 不是按字典序排列的，所以答案是 1。
示例 2：

输入：strs = ["xc","yb","za"]
输出：0
解释：
strs 的列已经是按字典序排列了，所以我们不需要删除任何东西。
注意 strs 的行不需要按字典序排列。
也就是说，strs[0][0] <= strs[0][1] <= ... 不一定成立。
示例 3：

输入：strs = ["zyx","wvu","tsr"]
输出：3
解释：
我们必须删掉每一列。


提示：

n == strs.length
1 <= n <= 100
1 <= strs[i].length <= 100
strs[i] 由小写英文字母组成

 */

/*
这是一道，很好的面试题，被低估了。
1876 的难度分数， 这题不是白给的！

// 这道题好难啊，
// 想到了一个好方法， 用 cmp 数组，来保持字典序。

算是入道了，以后坚持住，就这么来。
https://leetcode.cn/problems/delete-columns-to-make-sorted-ii/solutions/2598319/wei-hu-zi-dian-xu-go-by-zhang731-ux2j/
 */
func minDeletionSize(strs []string) int {
	n := len(strs[0])
	cmp := make([]int, len(strs)) // 0 equal, -1 < ,  1 >    // and cmp's value can not be -1, this is loop invariant we want to keep
	ans := 0
	for i := 0; i < n; i++ {
		cnt := 0
		ccmp := make([]int, len(strs))
		copy(ccmp, cmp) // save the previous ordering, if we need to delete the ith char, then we restore the samed ordering
		for j := 1; j < len(strs); j++ {
			if strs[j-1][i] < strs[j][i] {
				cmp[j] = 1
			} else if strs[j-1][i] > strs[j][i] {
				if cmp[j] != 1 {
					cnt++
					// keep cmp[j] the same, because ith char will be deleted
				}
				// else can do nothing
			} else if strs[j-1][i] == strs[j][i] {
				// 比较难处理
				// 不用处理。 哈哈, 因为 ordering 没有变化
			}
		}
		if cnt > 0 {
			ans++
			copy(cmp, ccmp) // since ith char will be deleted, we restore the same ordering that's [0:i-1] ordering of strs
		}
	}
	return ans
}

func minDeletionSize(strs []string) int {
	n := len(strs[0])
	cmp := make([]int, len(strs)) // 0 equal, -1 < ,  1 >    // and cmp's value can not be -1, this is loop invariant we want to keep
	ans := 0
	for i := 0; i < n; i++ {
		ccmp := make([]int, len(strs))
		copy(ccmp, cmp) // save the previous ordering
		for j := 1; j < len(strs); j++ {
			if strs[j-1][i] < strs[j][i] {
				cmp[j] = 1
			} else if strs[j-1][i] > strs[j][i] && cmp[j] != 1 {
				ans++
				copy(cmp, ccmp) // delete the ith char and restore the saved ordering that's [0:i-1] ordering of strs
				break
			}
		}
	}
	return ans
}

func minDeletionSize(strs []string) int {
	n := len(strs[0])
	cmp := make([]int, len(strs))
	ans := 0
	for i := 0; i < n; i++ {
		ccmp := make([]int, len(strs))
		copy(ccmp, cmp)
		for j := 1; j < len(strs); j++ {
			if strs[j-1][i] < strs[j][i] {
				cmp[j] = 1
			} else if strs[j-1][i] > strs[j][i] && cmp[j] != 1 {
				ans++
				copy(cmp, ccmp)
				break
			}
		}
	}
	return ans
}
