package Trie

type Trie struct {
	children [26]*Trie
	word     bool
	cnt      int
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	r := this
	for i := 0; i < len(word); i++ {
		a := int(word[i] - 'a')
		if r.children[a] == nil {
			r.children[a] = &Trie{}
		}
		r = r.children[a]
	}
	r.word = true
	r.cnt++
}

/***
看清楚啊， Search 和 StartWith 的 逻辑还是有区别的。
 */

func (this *Trie) Search(word string) bool {
	r := this
	for i := 0; i < len(word); i++ {
		a := int(word[i] - 'a')
		if r.children[a] == nil {
			return false
		}
		r = r.children[a]
	}
	return r.word
}

func (this *Trie) StartsWith(prefix string) bool {
	r := this
	for i := 0; i < len(prefix); i++ {
		a := int(prefix[i] - 'a')
		if r.children[a] == nil {
			return false
		}
		r = r.children[a]
	}
	return true
}
