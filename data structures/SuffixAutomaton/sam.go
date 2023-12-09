package SuffixAutomaton

/*
 Suffix Automaton 的实现

SAM 的其它应用还是不了解。

应用：
1. substring 的个数
LC 1698 此代码是 1698 的解

 */
type next [26]*node

type node struct {
	fa  *node
	ch  next
	len int
}

type sam struct {
	nodes []*node
	last  *node
}

func newSam() *sam {
	m := &sam{}
	m.last = m.newNode(nil, next{}, 0)
	return m
}

func (m *sam) newNode(fa *node, ch next, length int) *node {
	o := &node{fa: fa, ch: ch, len: length}
	m.nodes = append(m.nodes, o)
	return o
}

func (m *sam) append(c int) {
	last := m.newNode(m.nodes[0], next{}, m.last.len+1)
	for o := m.last; o != nil; o = o.fa {
		p := o.ch[c]
		if p == nil {
			o.ch[c] = last
			continue
		}
		if o.len+1 == p.len {
			last.fa = p
		} else {
			np := m.newNode(p.fa, p.ch, o.len+1)
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

func (m *sam) buildSam(s string) {
	for _, b := range s {
		m.append(int(b - 'a'))
	}
}

func countDistinct(s string) (ans int) {
	sam := newSam()
	sam.buildSam(s)
	for _, o := range sam.nodes[1:] {
		ans += o.len - o.fa.len
	}
	return
}
