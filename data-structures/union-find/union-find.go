package union_find

type UnionFind struct {
	fa       []int
	sz       []int
	comp_cnt int
}

func NewUnionFind(n int) *UnionFind {
	fa := make([]int, n)
	sz := make([]int, n)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	return &UnionFind{fa, sz, n}
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

func (u UnionFind) connected(x, y int) bool {
	x = u.find(x)
	y = u.find(y)
	return x == y
}
