package SuffixAutomaton

/*
LC1698 solution

code borrowed from https://oi-wiki.org/string/sam/#_15

 */
type state struct {
	len, link int
	next      map[byte]int
}

const MAXLEN = 10000

type sam struct {
	st       [MAXLEN * 2]state
	sz, last int
}

func (m *sam) sam_init() {
	m.st[0].len = 0
	m.st[0].link = -1
	m.st[0].next = make(map[byte]int)
	m.sz++
	m.last = 0
}

func (m *sam) sam_extend(c byte) {
	cur := m.sz
	m.sz++
	m.st[cur].next = make(map[byte]int)
	m.st[cur].len = m.st[m.last].len + 1
	p := m.last

	for p != -1 {
		if _, ok := m.st[p].next[c]; !ok {
			m.st[p].next[c] = cur
			p = m.st[p].link
		} else {
			break
		}
	}

	if p == - 1 {
		m.st[cur].link = 0
	} else {
		q := m.st[p].next[c]
		if m.st[p].len+1 == m.st[q].len {
			m.st[cur].link = q
		} else {
			clone := m.sz
			m.sz++
			m.st[clone].next = make(map[byte]int)
			m.st[clone].len = m.st[p].len + 1
			//st[clone].next = st[q].next
			for k, v := range m.st[q].next {
				m.st[clone].next[k] = v
			}
			m.st[clone].link = m.st[q].link

			for p != -1 && m.st[p].next[c] == q {
				m.st[p].next[c] = clone
				p = m.st[p].link
			}

			m.st[q].link = clone
			m.st[cur].link = clone
		}
	}
	m.last = cur
}

func (m *sam) buildSam(s string) {
	m.sam_init()
	for _, b := range s {
		m.sam_extend(byte(b))
	}
}
func countDistinct(s string) (ans int) {
	m := &sam{}
	m.buildSam(s)

	for i := 1; i < m.sz; i++ {
		ans += m.st[i].len - m.st[m.st[i].link].len
	}
	return
}
