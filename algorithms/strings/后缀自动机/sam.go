package sam

/*
后缀自动机 Suffix automaton (SAM)
将字符串的所有子串压缩后的结果

https://codeforces.com/blog/entry/20861 : 图很精彩，构造SAM的过程理解有障碍。
https://yutong.site/sam/


代码来自灵神：
https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/sam.go

代码的解释来自这里： https://oi-wiki.org/string/sam/
 */

type next [26]*node // 不定义成map，就是因为数组复制方便。

type node struct {
	//      len 为该节点（endpos 等价类）中最长的子串长度
	// fa.len+1 为该节点（endpos 等价类）中最短的子串长度
	// 等价类大小为 len-fa.len
	// 等价类中的每个子串都是其最长子串的一系列连续后缀，即长度组成了区间 [fa.len+1,len]
	// 这一系列连续后缀之后更短的后缀，要去 fa, fa.fa, fa.fa.fa, ... 上找
	fa  *node
	ch  next
	len int
	i   int
}

type sam struct {
	nodes []*node
	last  *node
}

func newSam(s string) *sam {
	m := &sam{}
	m.last = m.newNode(nil, next{}, 0, -1)
	m.buildSam(s)
	return m
}

func (m *sam) newNode(fa *node, _ch next, length, i int) *node {
	o := &node{fa: fa, ch: _ch, len: length, i: i}
	m.nodes = append(m.nodes, o)
	return o
}
func (m *sam) ord(c byte) byte { return c - 'a' }
func (m *sam) buildSam(s string) {
	for i, b := range s {
		m.append(i, int(m.ord(byte(b))))
	}
}

func (m *sam) append(i, c int) {
	last := m.newNode(m.nodes[0], next{}, m.last.len+1, i)

	for o := m.last; o != nil; o = o.fa {
		p := o.ch[c]
		if p == nil {
			o.ch[c] = last
			continue
		}

		if o.len+1 == p.len {
			last.fa = p
		} else {
			np := m.newNode(p.fa, p.ch, o.len+1, p.i)
			p.fa = np
			for ; o != nil && o.ch[c] == p; o = o.fa {
				o.ch[c] = np
			}
			last.fa = np
		}
		break
	}
	m.last = last
}
