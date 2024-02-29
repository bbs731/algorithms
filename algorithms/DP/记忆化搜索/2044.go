package dp

func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	mo := 0
	for _, n := range nums {
		mo |= n
	}

	ans := 0
	var dfs func(int, int)
	dfs = func(i, bitmask int) {
		if i == n {
			if bitmask == mo {
				ans++
			}
			return
		}
		dfs(i+1, bitmask)
		dfs(i+1, bitmask|nums[i])
	}

	dfs(0, 0)
	return ans
}

/****
怎么转换成 loop 的模式。
如果用for loop 来模拟所有的子集 ？


灵神是给过答案的：
https://leetcode.cn/problems/count-number-of-maximum-bitwise-or-subsets/solutions/1051495/o2n-zi-ji-huo-xie-fa-by-endlesscheng-bczp/

 */
