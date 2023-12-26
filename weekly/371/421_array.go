package weekly

/*
	感觉 array 的版本， 比 node pointer 的版本更加的好写。 两种都适用一下， 包括后期需要添加删除操作， 看看那种实现更舒服一些
 */
const MAXN = 3200000
const trieBitLen = 31 //30 for 1e9, 63 for int64, or bits.Len(MAX_VAL)

var next [MAXN][2]int
var num [MAXN]int
var cnt int // cnt 用 0 或者 1 做 root 都可以。

func insert(n int) {
	cur := 1
	for i := trieBitLen; i >= 0; i-- {
		bit := n >> i & 1
		if next[cur][bit] == 0 {
			cnt++
			next[cur][bit] = cnt
		}
		cur = next[cur][bit]
	}
	num[cur] = n
}

func find_maxXor(x int) int {
	cur := 1 // root
	for i := trieBitLen; i >= 0; i-- {
		bit := x >> i & 1
		if next[cur][bit^1] != 0 {
			cur = next[cur][bit^1]
		} else {
			cur = next[cur][bit]
		}
	}
	return x ^ num[cur]
}

func findMaximumXOR(nums []int) int {
	// leet code need to reset global variables
	cnt = 1
	next = [MAXN][2]int{}
	num = [MAXN]int{}

	ans := 0
	for _, x := range nums {
		insert(x)
		ans = max(ans, find_maxXor(x))
	}
	return ans
}
