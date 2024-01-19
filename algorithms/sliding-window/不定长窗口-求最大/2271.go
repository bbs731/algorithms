package weekly

import "sort"

/***

排序后，我们可以枚举毯子的摆放位置，然后计算毯子能覆盖多少块瓷砖。
实际上，毯子右端点放在一段瓷砖中间，是不如直接放在这段瓷砖右端点的（因为从中间向右移动，能覆盖的瓷砖数不会减少），所以可以枚举每段瓷砖的右端点来摆放毯子的右端点。

这样就可以双指针了，左指针 left\textit{left}left 需要满足其指向的那段瓷砖的右端点被毯子覆盖。

 */
func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	sort.Slice(tiles, func(i, j int) bool { return tiles[i][0] < tiles[j][0] })

	ans := 0
	left := 0
	cover := 0
	// 因为，枚举右端点， 写 sliding windows 写程序最方便， 线段树的方法也证明了， align z在最左端点，铺地毯，也是正确的。
	// 还是那个问题， 为什么是 tile 的右端点？ 为什么线段树的方法，左右端点都可以？
	for _, t := range tiles {
		tl, tr := t[0], t[1]
		cover += tr - tl + 1

		for tiles[left][1] < tr-carpetLen+1 {
			cover -= tiles[left][1] - tiles[left][0] + 1
			left++
		}

		if tr-carpetLen+1 >= tiles[left][0] && tr-carpetLen+1 <= tiles[left][1] { // 后面这个条件  tr-carpetLen+1 <= tiles[left][1]其实是不用的。 因为上面的 Loop 已经保证了。
			ans = max(ans, cover-(tr-carpetLen+1-tiles[left][0]))
		} else {
			ans = max(ans, cover)
		}

		//ans = max(ans, cover)
		//ans = max(ans, cover-max(tr-carpetLen+1-tiles[left][0], 0))
	}

	return ans

}
