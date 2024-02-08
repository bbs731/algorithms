package cache

/***
https://leetcode.cn/problems/lru-cache/solutions/2456294/tu-jie-yi-zhang-tu-miao-dong-lrupythonja-czgt/


请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

示例：
输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4

1 <= capacity <= 3000
0 <= key <= 10000
0 <= value <= 10^5
最多调用 2 * 105 次 get 和 put
 */

type LRUCache struct {
	table      map[int]int
	l          [][2]int
	start, end int
	capacity   int
	tick       int
	ticks      map[int]int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		table:    make(map[int]int, capacity),
		l:        make([][2]int, 0, capacity),
		capacity: capacity,
		tick:     0,
		ticks:    make(map[int]int, capacity),
	}
}

func (this *LRUCache) tickle(key int) {
	this.tick++
	this.l = append(this.l, [2]int{key, this.tick})
	this.ticks[key] = this.tick
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.table[key]; !ok {
		return -1
	} else {
		this.tickle(key)
		return v
	}
}

func (this *LRUCache) Put(key int, value int) {
	// 这里有个逻辑错误， 如果 key 使用来更新的怎么办？
	if _, ok := this.table[key]; ok {
		//// 更新操作
		this.tickle(key)
		this.table[key] = value
		return
	}

	// insert 操作。
	if len(this.table) == this.capacity {
		// need to pop one
		for len(this.l) > 0 && this.l[0][1] != this.ticks[this.l[0][0]] {
			this.l = this.l[1:]
		}
		old := this.l[0][0]
		delete(this.ticks, old)
		delete(this.table, old)
		this.l = this.l[1:]
	}

	this.tickle(key)
	this.table[key] = value
}
