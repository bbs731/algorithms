package weekly

/***


一位老师正在出一场由 n 道判断题构成的考试，每道题的答案为 true （用 'T' 表示）或者 false （用 'F' 表示）。老师想增加学生对自己做出答案的不确定性，方法是 最大化 有 连续相同 结果的题数。（也就是连续出现 true 或者连续出现 false）。

给你一个字符串 answerKey ，其中 answerKey[i] 是第 i 个问题的正确结果。除此以外，还给你一个整数 k ，表示你能进行以下操作的最多次数：

每次操作中，将问题的正确答案改为 'T' 或者 'F' （也就是将 answerKey[i] 改为 'T' 或者 'F' ）。
请你返回在不超过 k 次操作的情况下，最大 连续 'T' 或者 'F' 的数目。



示例 1：

输入：answerKey = "TTFF", k = 2
输出：4
解释：我们可以将两个 'F' 都变为 'T' ，得到 answerKey = "TTTT" 。
总共有四个连续的 'T' 。
示例 2：

输入：answerKey = "TFFT", k = 1
输出：3
解释：我们可以将最前面的 'T' 换成 'F' ，得到 answerKey = "FFFT" 。
或者，我们可以将第二个 'T' 换成 'F' ，得到 answerKey = "TFFF" 。
两种情况下，都有三个连续的 'F' 。
示例 3：

输入：answerKey = "TTFTTFTT", k = 1
输出：5
解释：我们可以将第一个 'F' 换成 'T' ，得到 answerKey = "TTTTTFTT" 。
或者我们可以将第二个 'F' 换成 'T' ，得到 answerKey = "TTFTTTTT" 。
两种情况下，都有五个连续的 'T' 。

 */

// 好题，标准的面试题目。
// 这道题，是 1004 的变种题目。 1004 做的很好，这道题，就不会了。

// 这个是正确的 O(n) 的解法。
func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(longestOnes(answerKey, k, 'T'), longestOnes(answerKey, k, 'F'))
}

func longestOnes(nums string, k int, zero byte) int {
	//n := len(nums)
	ans := 0
	l := 0
	cntZero := 0

	for i, n := range nums {
		if n == rune(zero) {
			cntZero++
		}

		for cntZero > k {
			// move left to satisfy constrain
			if nums[l] == zero {
				cntZero--
			}
			l++
		}
		ans = max(ans, i-l+1)
	}
	return ans
}

// 下面是错误的做法 是 O(n^2) 的做法， 有 O(n) sliding window 的做法。
func maxConsecutiveAnswers(answerKey string, k int) int {

	// 这里就不明智了， 聚合了之后， 还特别的不好办，使用 sliding window 会特别的麻烦。
	// record 连续的 element 的个数。
	cnts := []int{0, 1}
	pos := 1
	for i := 1; i < len(answerKey); i++ {
		s := answerKey[i]
		if s == answerKey[i-1] {
			cnts[pos]++
		} else {
			cnts = append(cnts, 1)
			pos++
		}
	}
	cnts = append(cnts, 0, 0)
	ans := 0
	// parse 一遍 cnts

	// 这里的复杂度是 n^2 如果 test case rigor 过不去。
	for i := 0; i < len(cnts); i++ {
		sum := 0
		credits := k
		for j := i + 1; j < len(cnts); j = j + 2 {
			sum += cnts[j-1]
			if credits < cnts[j] {
				sum += credits
				credits = 0
				break
			}
			credits -= cnts[j]
			sum += cnts[j]
		}
		ans = max(ans, sum+credits)
	}
	return min(ans, len(answerKey))
}
