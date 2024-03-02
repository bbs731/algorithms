package union_find

type UnionFind struct {
	fa       []int
	sz       []int
}

func NewUnionFind(n int) *UnionFind {
	fa := make([]int, n)
	sz := make([]int, n)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	return &UnionFind{fa, sz}
}

func (u UnionFind) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u UnionFind) unite(x, y int) {
	x = u.find(x)
	y = u.find(y)
	if x == y {
		return
	}
	if u.sz[x] < u.sz[y] {
		x, y = y, x
	}
	u.fa[y] = x
	u.sz[x] += u.sz[y]
}

func (u *UnionFind) getSize(x int)int {
	return u.sz[u.find(x)]
}


func reachableNodes(n int, edges [][]int, restricted []int) int {
	rm := make(map[int]bool, n)
	for _, r := range restricted {
		rm[r]= true
	}

	u := NewUnionFind(n)

	for _, e := range edges {
		from, to := e[0], e[1]
		if rm[from] || rm[to]{
			continue
		}
		u.unite(from, to)
	}

	return u.getSize(u.find(0))
}
