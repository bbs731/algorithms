package star

/*
按照边，来存贮图的一种方式。
这是最省空间的一种图的存储方式，（比 Adjacency list 的存储还要节省空间， 因为只用到了数组。 比 Adjacency list 的vector/list 还要节省空间， 对超大图的存储有用）
 */

const MAXM = 10000005
const MAXN = 10005

type Edge struct {
	to   int
	w    int
	next int
}

var edge [MAXM]Edge
var head [MAXN]int
var cnt int // 边的号码
var n, m int

func init() {

	for i := 0; i < n; i++ {
		head[i] = -1
	}
}

func add_edge(u, v, w int) {
	edge[cnt].to = v
	edge[cnt].w = w
	edge[cnt].next = head[u]
	head[u] = cnt
	cnt++
}

func loop() {

	u := 1
	for j := head[u]; j != -1; j = edge[j].next {

	}
}
