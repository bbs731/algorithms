package weekly

/*
数组中找两个树 异或值最大， 经典的 0-1 trie 模板题目。
代码还是参考灵神的吧， 为了以后好记，好找， 统一
 */

const trieBitLen = 31 //30 for 1e9, 63 for int64, or bits.Len(MAX_VAL)

type trieNode struct {
	son [2]*trieNode
	cnt int
}

type trie struct{ root *trieNode }

func (t *trie) insert(v int) *trieNode {
	o := t.root
	for i := trieBitLen; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b] == nil {
			o.son[b] = &trieNode{}
		}
		o = o.son[b]
	}
	return o
}

func (t *trie) maxXor(val int) int {
	o := t.root
	ans := 0

	for i := trieBitLen; i >= 0; i-- {
		b := val >> i & 1
		if o.son[b^1] != nil {
			ans |= 1 << i
			b ^= 1
		}
		o = o.son[b]
	}
	return ans
}

func findMaximumXOR(nums []int) int {
	root := &trie{&trieNode{}}
	ans := 0

	for _, x := range nums {
		root.insert(x)
		ans = max(ans, root.maxXor(x))
	}
	return ans
}
