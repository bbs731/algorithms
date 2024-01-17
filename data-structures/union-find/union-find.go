package leetcode

// recursive 的版本，是路径全压缩的！ loop 版本的 find 不是。
// 注意了 parent[x] = x 的初始化很重要， 要不然在创建 parent 的时候初始化， 要不然， 在 find 里面初始化 parent[x]
func main() {
	/*
			union-find template
	 */
	var n int
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int

	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(from, to int) {
		x, y := find(from), find(to)
		parent[x] = y
	}

}

// another implementation example

func equationsPossible(equations []string) {
	parent := make([]int, 26)
	for i := 0; i < 26; i++ {
		parent[i] = i
	}
	//....
}

func union(parent []int, index1, index2 int) {
	parent[find(parent, index1)] = find(parent, index2)
}

func find(parent []int, index int) int {
	for parent[index] != index {
		parent[index] = parent[parent[index]]
		index = parent[index]
	}
	return index
}
