package binary_search

import "sort"

/****

给你一个整数数组 heights ，表示建筑物的高度。另有一些砖块 bricks 和梯子 ladders 。

你从建筑物 0 开始旅程，不断向后面的建筑物移动，期间可能会用到砖块或梯子。

当从建筑物 i 移动到建筑物 i+1（下标 从 0 开始 ）时：

如果当前建筑物的高度 大于或等于 下一建筑物的高度，则不需要梯子或砖块
如果当前建筑的高度 小于 下一个建筑的高度，您可以使用 一架梯子 或 (h[i+1] - h[i]) 个砖块
如果以最佳方式使用给定的梯子和砖块，返回你可以到达的最远建筑物的下标（下标 从 0 开始 ）。

输入：heights = [4,2,7,6,9,14,12], bricks = 5, ladders = 1
输出：4
解释：从建筑物 0 出发，你可以按此方案完成旅程：
- 不使用砖块或梯子到达建筑物 1 ，因为 4 >= 2
- 使用 5 个砖块到达建筑物 2 。你必须使用砖块或梯子，因为 2 < 7
- 不使用砖块或梯子到达建筑物 3 ，因为 7 >= 6
- 使用唯一的梯子到达建筑物 4 。你必须使用砖块或梯子，因为 6 < 9
无法越过建筑物 4 ，因为没有更多砖块或梯子。
示例 2：

输入：heights = [4,12,2,7,3,18,20,3,19], bricks = 10, ladders = 2
输出：7
示例 3：

输入：heights = [14,3,19,3], bricks = 17, ladders = 0
输出：3


提示：

1 <= heights.length <= 10^5
1 <= heights[i] <= 10^6
0 <= bricks <= 10^9
0 <= ladders <= heights.length

 */

// 1962 的难度，感觉怎么样？
func furthestBuilding(heights []int, bricks int, ladders int) int {
	n := len(heights)

	// now 二分答案
	l, r := -1, n-1 // (l, r]  // 这个区间取的，左开右闭，写的太漂亮了！
	// 先 true 后 false 的情况。 是不是 左开右闭的区间，最好呢？ 应为是找，最右边的 true 的位置。
	for l < r {
		mid := (l + r + 1) >> 1
		cost := make([]int, 0, n)

		for i := 0; i <= mid; i++ {
			j := i + 1
			if j <= mid && heights[j] > heights[i] {
				cost = append(cost, heights[j]-heights[i])
			}
		}
		sort.Ints(cost) //这个竟然忘了！
		end := len(cost)
		if ladders != 0 {
			end -= ladders
		}
		sum := 0
		for j := 0; j < end; j++ {
			sum += cost[j]
		}

		if sum <= bricks {
			l = mid
		} else {
			r = mid - 1
		}
	}
	// l == r
	return l
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	n := len(heights)

	// 先 true 后 false 的情况。 是不是 左开右闭的区间，最好呢？ 应为是找，最右边的 true 的位置。
	// 想把 技巧一 和 技巧二融合也不是不能，但是难度系数挺高啊，不如上面的 （ ] 区间写的简单。
	return -1 + sort.Search(n, func(x int) bool { // 这样写是不是清楚一些， Search 的值域是 [0, n]所求值域为 [-1,n-1]
		// x += l
		x++ // 这里根据  // sort.Search 的使用技巧·其一
		cost := make([]int, 0, n)
		for i := 0; i < x; i++ {
			j := i + 1
			if j < x && heights[j] > heights[i] {
				cost = append(cost, heights[j]-heights[i])
			}
		}
		sort.Ints(cost) //这个竟然忘了！
		end := len(cost)
		if ladders != 0 {
			end -= ladders
		}
		sum := 0
		for j := 0; j < end; j++ {
			sum += cost[j]
		}
		//sum <= bricks  这里取反
		return sum > bricks
	}) // - 1 // 灵神来解释一下这里！ (这道题，灵神没有给过答案）

	// 最后返回的值域应该是 [-1, n-1]
}

// 先准备，在根据准备的数据去测， 这样，逻辑太复杂了。 所以，思路，正的，和反的， 写出来难度系数不一样啊！找简单的方法去写, 不出错。
func furthestBuilding(heights []int, bricks int, ladders int) int {
	n := len(heights)
	climeb := make([]int, 0)
	cross := make([]int, 0)

	// 这个 Prepare 的逻辑，是不是写的太复杂了？ 有更简单的解法吗？ 好像有啊！哈哈， 不用事先准备好啊，可以安装题目的要求现求也是可以的。
	for i := 0; i < n; i++ {
		start := i
		next := start + 1
		if next < n && heights[next] <= heights[start] {
			for ; next < n && heights[next] <= heights[next-1]; next++ {
			}
		} else {
			next = start
		}

		if next+1 < n {
			climeb = append(climeb, heights[next+1]-heights[next])
			// [start, nex+1]
			if len(cross) == 0 {
				cross = append(cross, next+1-start)
			} else {
				// accumulate
				cross = append(cross, cross[len(cross)-1]+next+1-start)
			}
		} else {
			// now next is the last element  next = n-1
			// 处理最后一个 block
			if len(cross) != 0 {
				cross[len(cross)-1] += next - start
			}
		}
		i = next
	}

	if len(climeb) == 0 {
		return n - 1
	}

	// now 二分答案

	l, r := -1, len(climeb)
	// 先 true 后 false 的情况。
	for l+1 < r {
		mid := (l + r) >> 1

		cost := make([]int, 0, len(climeb))
		copy(cost, climeb[:mid+1])
		sort.Ints(cost)
		sum := 0
		for i := 0; i <= mid-ladders; i++ {
			sum += cost[i]
		}

		if ladders > mid || sum <= bricks {
			l = mid
		} else {
			r = mid
		}
	}
	// l + 1 == r
	if l == -1 {
		return 0
	}
	return cross[l]
}
