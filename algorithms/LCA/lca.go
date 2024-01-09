package LCA

const MAXN = 1 << 14

type edge struct {
	to, wt int
}

var (
	Log2   [MAXN]int
	fa     [MAXN][20]int
	weight [MAXN][20][26]int
	dep    [MAXN]int
	//vis  [MAXN]bool  // 如果 dfs 传入 fath, 因为是树，所以可以不用 vis 数组。但是输入，有可能不告诉你根节点是哪个。
	g [][]edge
)

func dfs(cur, fath, d int) {
	//if vis[cur] {
	//	return
	//}
	//vis[cur] = true
	//dep[cur] = dep[fath] + 1  // 因为 root 的 parent -1 , 不能当做 index 所以显示的 pass in d as depth
	dep[cur] = d
	fa[cur][0] = fath

	// 这个可以写在外层吗？  答案是可以的。看下面 loop 形式倍增的模板
	for i := 1; i <= Log2[dep[cur]]; i++ {
		fa[cur][i] = fa[fa[cur][i-1]][i-1]
	}
	for _, e := range g[cur] {
		if w := e.to; w != fath {
			weight[w][0][e.wt] = 1
			dfs(e.to, cur, d+1)
		}
	}
}

// 这里写的思路好清晰
func lca(a, b int) int {
	if dep[a] > dep[b] {
		a, b = b, a
	}

	for dep[a] != dep[b] {
		b = fa[b][Log2[dep[b]-dep[a]]]
	}
	if a == b {
		return a
	}

	for k := Log2[dep[a]]; k >= 0; k-- {
		if fa[a][k] != fa[b][k] {
			a, b = fa[a][k], fa[b][k]
		}
	}
	return fa[a][0]
}

func main() {
	var edges [][]int
	var n int // 节点是从 0 到 n-1
	for i := 2; i < n; i++ {
		Log2[i] = Log2[i/2] + 1
	}
	// need to build graph
	g := make([][]edge, n)
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]-1
		g[v] = append(g[v], edge{w, wt})
		g[w] = append(g[w], edge{v, wt})
	}

	dfs(0, -1, 0)

	// 这段倍增的模板是，写在外面（现对于写在 dfs 里面）的方式。 是可以的。 但是需要注意 Loop的顺序，需要先loop k 再 loop node number
	// 这段模板，需要写在 dfs 之后，在计算完 depth 和 第一层 fa[node][0] 之后。
	for k := 1; k <= Log2[n-1]; k++ {
		for i := 0; i < n; i++ {
			if fa[i][k-1] != -1 {
				fa[i][k] = fa[fa[i][k-1]][k-1]
			}
		}
	}

	return
}
