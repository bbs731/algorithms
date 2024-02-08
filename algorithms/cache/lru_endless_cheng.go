package cache

import "container/list"

/****
灵神的实现，太牛了！ 之前没使用过 golang 的 list container 没想到有这么简洁的用法。
 */
type entry struct {
	key, value int
}

type LRUCache struct {
	capacity  int
	list      *list.List // 双向链表
	keyToNode map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, list.New(), map[int]*list.Element{}}
}

func (c *LRUCache) Get(key int) int {
	node := c.keyToNode[key]
	if node == nil { // 没有这本书
		return -1
	}
	c.list.MoveToFront(node) // 把这本书放在最上面
	return node.Value.(entry).value
}

func (c *LRUCache) Put(key, value int) {
	if node := c.keyToNode[key]; node != nil { // 有这本书
		node.Value = entry{key, value} // 更新
		c.list.MoveToFront(node)       // 把这本书放在最上面
		return
	}
	c.keyToNode[key] = c.list.PushFront(entry{key, value}) // 新书，放在最上面
	if len(c.keyToNode) > c.capacity { // 书太多了
		delete(c.keyToNode, c.list.Remove(c.list.Back()).(entry).key) // 去掉最后一本书
	}
}
