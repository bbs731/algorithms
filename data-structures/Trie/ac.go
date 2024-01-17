package Trie

/*

leetcoce: 1032

// EXTRA: AC 自动机 Aho–Corasick algorithm / Deterministic Finite Automaton (DFA)
// https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm
// https://en.wikipedia.org/wiki/Deterministic_finite_automaton
// 基础实现 https://zhuanlan.zhihu.com/p/80325757
// 基础实现 https://www.cnblogs.com/nullzx/p/7499397.html
// 改进实现 https://oi-wiki.org/string/ac-automaton/


https://oi-wiki.org/string/ac-automaton/
https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/trie.go
 */

// 之前写给 feedback-details-go 项目的 AC 自动机， 截取了一些和业务无关的代码。
// 可能是不完备的，仅当例子，不做为 template 代码。 // 以后做题的时候，完善加上 AC 通用的模板代码

type node struct {
	son        [128]*node
	fail       *node
	patternIdx int
	cnt        int //（子树中）完整字符串的个数
}

type Ac struct {
	root     *node
	patterns []string
	nodeCnt  int
}

// trie insert
func (t *Ac) put(word string, idx int) {
	o := t.root

	for i, c := range word {
		if o.son[c] == nil {
			o.son[c] = &node{}
			t.nodeCnt++
		}
		o = o.son[c]
		//o.cnt++  写法一: 统计 o 对应的字符串是多少个完整字符串patterns 的前缀。
	}
	o.patternIdx = idx
	o.cnt++ // 写法二： 统计 o 上有多少个完成的字符串
}

func (t *Ac) buildDFA() {
	t.root.fail = t.root
	q := make([]*node, 0, t.nodeCnt)

	for i, son := range t.root.son[:] {
		if son == nil {
			t.root.son[i] = t.root
		} else {
			son.fail = t.root
			q = append(q, son)
		}
	}

	for len(q) > 0 {
		o := q[0]
		q = q[1:]

		if o.fail == nil {
			o.fail = t.root
		}

		for i, son := range o.son[:] {
			if son != nil {
				son.fail = o.fail.son[i]
				q = append(q, son)
			} else {
				o.son[i] = o.fail.son[i]
			}
		}
	}
}

// 返回一个 pos 列表，其中 pos[i] 表示 patterns[i] 的【首字母】在文本串 text 的所有位置（未找到时为空）
func (t *Ac) acMatch(text string, patters []string) [][]int {
	pos := make([][]int, len(t.patterns))
	o := t.root

	for i, b := range text {
		if b >= 128 {
			// invalid input char. continue or panic
			continue
		}

		o = o.son[b]
		if o == nil {
			o = t.root
			continue
		}
		for f := o; f != t.root; f = f.fail {
			if pid := f.patternIdx; pid != 0 {
				// pattern matched.
				pos[pid] = append(pos[pid], i-len(t.patterns[pid])+1)
			}
		}
	}
	return pos
}
