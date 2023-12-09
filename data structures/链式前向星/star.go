package star

/*
按照边，来存贮图的一种方式。
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
