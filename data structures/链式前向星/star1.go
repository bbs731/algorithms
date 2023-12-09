package star

var n, m, cnt int

var head [MAXN]int
var next [MAXM]int
var to [MAXM]int
var weights [MAXM]int

func init() {
	for i := 0; i < n; i++ {
		head[i] = -1
	}
}

func add_edge(u, v, w int) {
	to[cnt] = v
	weights[cnt] = w
	next[cnt] = head[u]
	head[u] = cnt
	cnt++
}

func loop() {
	u := 2

	for j := head[u]; j != -1; j = next[j] {

	}
}
