package weekly

//You are given four integers sx, sy, fx, fy, and a non-negative integer t.
//
// In an infinite 2D grid, you start at the cell (sx, sy). Each second, you must
// move to any of its adjacent cells.
//
// Return true if you can reach cell (fx, fy) after exactly t seconds, or false
//otherwise.
//
// A cell's adjacent cells are the 8 cells around it that share at least one cor
//ner with it. You can visit the same cell several times.
//
//
// Example 1:
//
//
//Input: sx = 2, sy = 4, fx = 7, fy = 7, t = 6
//Output: true
//Explanation: Starting at cell (2, 4), we can reach cell (7, 7) in exactly 6 se
//conds by going through the cells depicted in the picture above.
//
//
// Example 2:
//
//
//Input: sx = 3, sy = 1, fx = 7, fy = 3, t = 3
//Output: false
//Explanation: Starting at cell (3, 1), it takes at least 4 seconds to reach cel
//l (7, 3) by going through the cells depicted in the picture above. Hence, we can
//not reach cell (7, 3) at the third second.
//
//
//
// Constraints:
//
//
// 1 <= sx, sy, fx, fy <= 109
// 0 <= t <= 109
//

func abs(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}
func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	if sx == fx && sy == fy { // sy 总写错
		return t != 1 // 这里是个大坑。
	}
	return max(abs(sx, fx), abs(sy, fy)) <= t
}
