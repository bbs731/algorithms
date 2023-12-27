package _600_1900

/*
给你两个字符串数组 words1 和 words2。

现在，如果 b 中的每个字母都出现在 a 中，包括重复出现的字母，那么称字符串 b 是字符串 a 的 子集 。

例如，"wrr" 是 "warrior" 的子集，但不是 "world" 的子集。
如果对 words2 中的每一个单词 b，b 都是 a 的子集，那么我们称 words1 中的单词 a 是 通用单词 。

以数组形式返回 words1 中所有的通用单词。你可以按 任意顺序 返回答案。



示例 1：

输入：words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["e","o"]
输出：["facebook","google","leetcode"]
示例 2：

输入：words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["l","e"]
输出：["apple","google","leetcode"]
示例 3：

输入：words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["e","oo"]
输出：["facebook","google"]
示例 4：

输入：words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["lo","eo"]
输出：["google","leetcode"]
示例 5：

输入：words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["ec","oc","ceo"]
输出：["facebook","leetcode"]


提示：
1 <= words1.length, words2.length <= 10^4
1 <= words1[i].length, words2[i].length <= 10
words1[i] 和 words2[i] 仅由小写英文字母组成
words1 中的所有字符串 互不相同
 */

/*
这就是 枚举的技巧： O（2 * 26 * 10^4）
枚举的力量。枚举的力量。

难度分： 1624
正常的面试题应该就是类似的难度。

 */
func wordSubsets(words1 []string, words2 []string) []string {
	dict2 := make(map[int]int)

	for _, w := range words2 {
		cnt := [26]int{}
		for _, c := range w {
			cnt[c-'a']++
		}

		for i := 0; i < 26; i++ {
			dict2[i] = max(dict2[i], cnt[i])
		}
	}
	ans := make([]string, 0)

	for _, w := range words1 {
		pass := true
		cnt := [26]int{}
		for _, c := range w {
			cnt[c-'a']++
		}

		for i := 0; i < 26; i++ {
			if cnt[i] < dict2[i] {
				pass = false
				break
			}
		}
		if pass {
			ans = append(ans, w)
		}
	}
	return ans
}

/*
这样枚举会超时 O（ 10^8)
 */
func wordSubsets(words1 []string, words2 []string) []string {

	process := func(words []string) map[string]map[byte]int {
		dict := make(map[string]map[byte]int, len(words))
		for i := 0; i < len(words); i++ {
			cnt := make(map[byte]int)
			for _, c := range words[i] {
				cnt[byte(c)]++
			}
			dict[words1[i]] = cnt
		}
		return dict
	}

	d1 := process(words1)
	d2 := process(words2)
	//sort.Strings(words2)
	ans := make([]string, 0)

	for _, word1 := range words1 {
		dh := d1[word1]
		pass := true
		//for j := len(words2) - 1; j >= 0; j-- {
		for _, m := range d2 {
			m := d2[words2[j]]
			for k, v := range m {
				if cnt, ok := dh[k]; !ok || cnt < v {
					pass = false // failed early break
					break
				}
			}
			if pass == false {
				break
			}
		}
		if pass {
			ans = append(ans, word1)
		}
	}
	return ans
}
