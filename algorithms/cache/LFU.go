package cache

import "container/list"

/***
你虽然，做出来了， 但是，还是有些瑕疵。 最外层的 List 长度你是没做任何限制的 （这个内存和时间复杂度，将来有会有问题）
看看， 灵神的答案， 他没有用 list 套 list, 他用 map 套在了 list 外面。 这样比list(list) 更加容易的处理，中间freq 有断层的问题。

但是值得鼓励啊， 这么长的代码，和逻辑，基本上一次写对， 但是也花了一个小时
加油！
 */
type entry struct {
	key, value int
}

type LFUCache struct {
	capacity int
	list     *list.List // we create a list of list, bottom is the last list element of list, and bottom itself is a list
	//bottom    *list.List
	keyToNode map[int]*list.Element
	keyToList map[int]*list.Element
}

func Constructor(capacity int) LFUCache {
	l := list.New()
	bottom := list.New()
	l.PushFront(bottom)

	return LFUCache{capacity, l, map[int]*list.Element{}, map[int]*list.Element{}}

}

func (c *LFUCache) Get(key int) int {
	node := c.keyToNode[key]
	if node == nil {
		return -1
	}
	l := c.keyToList[key]
	v := c.keyToNode[key].Value.(entry).value

	var nl *list.List
	if l.Prev() == nil {
		nl = list.New()
		c.list.PushFront(nl)
	} else {
		nl = l.Prev().Value.(*list.List)
	}
	l.Value.(*list.List).Remove(node)
	c.keyToNode[key] = nl.PushFront(entry{key, v})
	c.keyToList[key] = l.Prev()

	return v
}

func (c *LFUCache) Put(key int, value int) {
	if node := c.keyToNode[key]; node != nil {
		// update an existing entry
		node.Value = entry{key, value}
		var nl *list.List
		l := c.keyToList[key]
		if l.Prev() == nil {
			nl = list.New()
			c.list.PushFront(nl)
		} else {
			nl = l.Prev().Value.(*list.List)
		}
		// remove from the old list
		_ = l.Value.(*list.List).Remove(node)
		// adding into new list
		c.keyToNode[key] = nl.PushFront(node.Value)
		c.keyToList[key] = l.Prev()
		return
	}
	// a new insertion

	if len(c.keyToNode) == c.capacity {
		// 这里不能删除， 删除会有错误。
		//for c.bottom.Len() == 0 {
		//	c.list.Remove(c.list.Back())
		//	c.bottom = c.list.Back().Value.(*list.List)
		//}
		l := c.list.Back()
		for l.Value.(*list.List).Len() == 0 {
			l = l.Prev()
		}
		//key := c.bottom.Remove(c.bottom.Back()).(entry).key
		ll := l.Value.(*list.List)
		key := ll.Remove(ll.Back()).(entry).key
		delete(c.keyToNode, key)
		delete(c.keyToList, key)
	}

	c.keyToNode[key] = c.list.Back().Value.(*list.List).PushFront(entry{key, value})
	c.keyToList[key] = c.list.Back() // which is bottom

}
