package dfs

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	f := make([]bool, n)

	var dfs func(int)
	dfs = func(i int) {
		// visited before
		if f[i] == true {
			return
		}

		f[i] = true
		for _, j := range rooms[i] {
			dfs(j)
		}
	}
	dfs(0)

	for i := 0; i < n; i++ {
		if f[i] == false {
			return false
		}
	}
	return true
}
