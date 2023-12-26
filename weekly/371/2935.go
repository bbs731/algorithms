package weekly

import "sort"

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
		o.cnt++
	}
	return o
}

func (t *trie) remove(v int) {
	o := t.root
	for i := trieBitLen; i >= 0; i-- {
		b := v >> i & 1
		o = o.son[b]
		o.cnt--
	}
}

func (t *trie) maxXor(val int) int {
	o := t.root
	ans := 0

	for i := trieBitLen; i >= 0; i-- {
		b := val >> i & 1
		if o.son[b^1] != nil && o.son[b^1].cnt != 0 {
			ans |= 1 << i
			b ^= 1
		}
		o = o.son[b]
	}
	return ans
}

func maximumStrongPairXor(nums []int) int {
	sort.Ints(nums)
	ans := 0
	//q := []int{} // 这里直接用一个 left 指针就可以。
	left := 0
	root := &trie{&trieNode{}}
	for _, y := range nums {
		root.insert(y)
		for nums[left]*2 < y {
			root.remove(nums[left])
			left++
		}
		//q = q[i:]
		ans = max(ans, root.maxXor(y))
	}
	return ans
}
