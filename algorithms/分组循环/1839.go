package loop



/****
https://leetcode.cn/problems/longest-substring-of-all-vowels-in-order/solutions/742586/mo-ni-by-endlesscheng-otyy/
灵神的答案，用上了 regular expression 了。

这个 分组循环好难。
我感觉我这个分组循环的逻辑写的比灵神本题的答案清晰。  这道分组循环的，多了一部分， 需要准备工作！
 */
func longestBeautifulSubstring(word string) int {
	n := len(word)
	ans := 0
	vowls := []byte{'a', 'e', 'i', 'o', 'u'}

	for i:=0; i <n; {
		//  这里需要先做准备工作， 找到第一个 a 的位置。
		if word[i] != 'a' {
			i++
			continue
		}
		start := i
		for v := 0; v<len(vowls); v++ {
			for ; i < n && word[i] == vowls[v]; i++ {
			}
			if word[i-1] != vowls[v] {
				break
			} else{
				if v == 4 {
					// record the answer
					ans = max(ans, i- start)
				}
			}
		}
	}
	return ans
}


