package dp

/*
看灵神的题解：
https://leetcode.cn/problems/maximize-palindrome-length-from-subsequences/solutions/2285215/shi-pin-qiao-miao-zhuan-huan-516-bian-xi-jhrt/
比自己的想法，简单了好多。 这道题是 516 的变形题， 看看 516 的模板，把 516 的解简化了之后，回来再看这道题。

func longestPalindrome(word1, word2 string) (ans int) {
    s := word1 + word2
    n := len(s)
    f := make([][]int, n)
    for i := range f {
        f[i] = make([]int, n)
    }
    for i := n - 1; i >= 0; i-- {
        f[i][i] = 1
        for j := i + 1; j < n; j++ {
            if s[i] == s[j] {
                f[i][j] = f[i+1][j-1] + 2
                if i < len(word1) && j >= len(word1) {
                    ans = max(ans, f[i][j]) // f[i][j] 一定包含 s[i] 和 s[j]
                }
            } else {
                f[i][j] = max(f[i+1][j], f[i][j-1])
            }
        }
    }
    return
}

func max(a, b int) int { if a < b { return b }; return a }

作者：灵茶山艾府
链接：https://leetcode.cn/problems/maximize-palindrome-length-from-subsequences/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

 */
/*
	这道题，相当于两道 区间 DP 的问题， 第一遍做预处理。 第二个 区间DP 问题，的边界条件，用到第一个 DP 的预处理结果。
	太酸爽了， 做了整整一个上午。

	看看灵神的答案， 把 word1 + word2 两个字符串想加，找子串的palindrome 是不是就豁然开朗了。
 */

/*
	dfs(i, j)  i start position from word1 [0.. n1-1],   j end position from word2 [0..n2-1]
	定义为， 能总成的最长 palindrome 子串的最大长度。

	dfs(i, j) = max { dfs(i+1,j-1) + 2 //if word1[i] == word2[j],  dfs(i+1,j-1) // 都不选， dfs(i+1, j), dfs(i, j-1) }


边界条件：
	i > n1  return 0 or j+1  word2[j] is palindrome
	j < 0 return 0 or n1- i + 1 word1[i:] is palindrome

	return  dfs(0, n2-1)

 */

/*

dfs(i, j) 定义为 s[i..j] 中 subsequence 可以形成 palindrome 的最大长度。
dfs(i,j) = max(dfs(i+1,j), dfs(i, j-1),   dfs(i+1, j-1) + 2 if s[i] == s[j])

dfs(i, i) = 1
dfs(i+1, i) = 0

f[i][j+1] = max(f[i+1][j+1] , f[i][j], f[i+1][j] +2)  // 其实，可以不做 j 下标的转换。
f[i]pj[ = max(f[i+1][j] , f[i][j-1], f[i+1][j-1] +2)
i 是倒序遍历， j 是正序遍历。

初始化条件：
f[i][i+1] = 1  // f[i][i] = 1
f[i+1][i+1] = 0  //f[i+1][i] = 0


看看上面灵神关于 f[i][j] 的计算， 有两点区别。
1. 没有做 j 的 index 变换  (变成 j+1)
2. 初始化更加的简洁。

 */
func maxSubsequencePalindrom(s string) [][]int {
	n := len(s)
	//inf := int(1e9)

	f := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		f[i] = make([]int, n)
		f[i][i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ { // j loop from i+1  这是个天坑， 卡在这里很长时间。
			//f[i][j+1] = -inf
			if s[i] == s[j] {
				f[i][j] = max(f[i][j], f[i+1][j-1]+2)
			} else {
				f[i][j] = max(f[i][j], f[i+1][j], f[i][j-1])
			}
		}
	}
	return f
}

func longestPalindrome(word1 string, word2 string) int {
	n1 := len(word1)
	n2 := len(word2)
	inf := int(1e9)

	var dfs func(int, int, int) int
	p1 := maxSubsequencePalindrom(word1)
	p2 := maxSubsequencePalindrom(word2)

	cache := make([][][]int, n1)
	for i := 0; i < n1; i++ {
		cache[i] = make([][]int, n2)
		for j := 0; j < n2; j++ {
			cache[i][j] = make([]int, 2)
			cache[i][j][0] = -1
			cache[i][j][1] = -1
		}
	}
	// 这个 valid 也是殚精竭虑啊， 如果出现这样的小技巧， 其实说明，题目还有更简单的解法。
	dfs = func(i, j int, valid int) int {
		if i >= n1 && j < 0 {
			if valid == 1 {
				return 0
			}
			return -inf
		}

		if i >= n1 {
			if valid == 1 {
				return p2[0][j]
			}
			return 0
		}

		if j < 0 {
			if valid == 1 {
				return p1[i][n1-1]
			}
			return 0
		}
		if cache[i][j][valid] != -1 {
			return cache[i][j][valid]
		}
		ans := -inf
		ans = max(ans, dfs(i+1, j, valid), dfs(i, j-1, valid), dfs(i+1, j-1, valid))
		if word1[i] == word2[j] {
			ans = max(ans, dfs(i+1, j-1, 1)+2)
		}
		cache[i][j][valid] = ans
		return ans
	}
	return dfs(0, n2-1, 0)
}

func longestPalindrome_wrong(word1 string, word2 string) int {
	n1 := len(word1)
	n2 := len(word2)
	inf := int(1e9)

	var dfs func(int, int, int) int

	// 预处理一下 palindrome
	// isP2[j] means from 0 to j of word2 is palindrome
	// isP1[i] means from i to n1 is palindrome

	f1 := make([][]bool, n1+1)
	for i := n1 - 1; i >= 0; i-- {
		f1[i] = make([]bool, n1+1)
		f1[i][i+1] = true
		if i+1 < n1 {
			f1[i+1][i+1] = true
		}
	}

	for i := n1 - 1; i >= 0; i-- {
		for j := i + 1; j < n1; j++ {
			if word1[i] == word1[j] && f1[i+1][j] {
				f1[i][j+1] = true
			}
		}
	}
	f1[0][n1] = false

	f2 := make([][]bool, n2+1)
	for i := n2 - 1; i >= 0; i-- {
		f2[i] = make([]bool, n2+1)
		f2[i][i+1] = true
		if i+1 < n2 {
			f2[i+1][i+1] = true
		}
	}
	for i := n2 - 1; i >= 0; i-- {
		for j := i + 1; j < n2; j++ {
			if word2[i] == word2[j] && f2[i+1][j] {
				f2[i][j+1] = true
			}
		}
	}
	f2[0][n2] = false

	//fmt.Println(f1)
	//fmt.Println(f2)

	cache := make([][][]int, n1)
	for i := 0; i < n1; i++ {
		cache[i] = make([][]int, n2)
		for j := 0; j < n2; j++ {
			cache[i][j] = make([]int, 2)
			cache[i][j][0] = -1
			cache[i][j][1] = -1
		}
	}
	dfs = func(i, j int, valid int) int {
		if i >= n1 && j < 0 {
			if valid == 1 {
				return 0
			}
			return -inf
		}

		if i >= n1 {

			// 这个是错误的， 我们需要计算的是，  word2[0, j] 这个字符串中，最大回文子串的长度。
			if f2[0][j+1] && valid == 1 {
				//fmt.Printf("i is :%d, j is : %d and return %d\n", i, j, j+1)
				return j + 1
			}
			if valid == 1 {
				return 1
			}
			return 0
		}

		if j < 0 {
			// 同理，在 valid 的情况下， 我们应该返回， word1[i,n1-1] 这个子串中，最大回文子串的长度。
			if f1[i][n1] && valid == 1 {
				//fmt.Printf("its i : %d\n", i)
				return n1 - i
			}
			if valid == 1 {
				return 1
			}
			return 0
		}
		if cache[i][j][valid] != -1 {
			return cache[i][j][valid]
		}
		ans := -inf
		ans = max(ans, dfs(i+1, j, valid), dfs(i, j-1, valid), dfs(i+1, j-1, valid))
		if word1[i] == word2[j] {
			ans = max(ans, dfs(i+1, j-1, 1)+2)
		}
		cache[i][j][valid] = ans
		return ans
	}
	return dfs(0, n2-1, 0)
}
