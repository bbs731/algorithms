package weekly

import "sort"

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	//n := len(tiles)
	//a := make([]int, 2*n*4)
	//lazyRoot.build(a, 1, 2*n)
	sort.Slice(tiles, func(i, j int) bool { return tiles[i][0] < tiles[j][0] })

	for _, t := range tiles {
		lazyRoot.update(t[0], t[1], 1)
	}

	ans := 0
	for _, t := range tiles {
		ans = max(ans, lazyRoot.query(t[0], t[0]+carpetLen-1))
	}
	return ans
}
