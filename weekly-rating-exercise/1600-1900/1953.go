package _600_1900

/*

给你 n 个项目，编号从 0 到 n - 1 。同时给你一个整数数组 milestones ，其中每个 milestones[i] 表示第 i 个项目中的阶段任务数量。

你可以按下面两个规则参与项目中的工作：

每周，你将会完成 某一个 项目中的 恰好一个 阶段任务。你每周都 必须 工作。
在 连续的 两周中，你 不能 参与并完成同一个项目中的两个阶段任务。
一旦所有项目中的全部阶段任务都完成，或者仅剩余一个阶段任务都会导致你违反上面的规则，那么你将 停止工作 。注意，由于这些条件的限制，你可能无法完成所有阶段任务。

返回在不违反上面规则的情况下你 最多 能工作多少周。



示例 1：

输入：milestones = [1,2,3]
输出：6
解释：一种可能的情形是：
​​​​- 第 1 周，你参与并完成项目 0 中的一个阶段任务。
- 第 2 周，你参与并完成项目 2 中的一个阶段任务。
- 第 3 周，你参与并完成项目 1 中的一个阶段任务。
- 第 4 周，你参与并完成项目 2 中的一个阶段任务。
- 第 5 周，你参与并完成项目 1 中的一个阶段任务。
- 第 6 周，你参与并完成项目 2 中的一个阶段任务。
总周数是 6 。
示例 2：

输入：milestones = [5,2,1]
输出：7
解释：一种可能的情形是：
- 第 1 周，你参与并完成项目 0 中的一个阶段任务。
- 第 2 周，你参与并完成项目 1 中的一个阶段任务。
- 第 3 周，你参与并完成项目 0 中的一个阶段任务。
- 第 4 周，你参与并完成项目 1 中的一个阶段任务。
- 第 5 周，你参与并完成项目 0 中的一个阶段任务。
- 第 6 周，你参与并完成项目 2 中的一个阶段任务。
- 第 7 周，你参与并完成项目 0 中的一个阶段任务。
总周数是 7 。
注意，你不能在第 8 周参与完成项目 0 中的最后一个阶段任务，因为这会违反规则。
因此，项目 0 中会有一个阶段任务维持未完成状态。


提示：

n == milestones.length
1 <= n <= 10^5
1 <= milestones[i] <= 10^9

 */

/*
1800 分的难度，开始搞不定了！

我的天啊！ 看答案， 如何证明贪心是正确的？
 */
func numberOfWeeks(m []int) int64 {
	longest := m[0]

	sum := 0
	for _, x := range m {
		sum += x
		if x > longest {
			longest = x
		}
	}
	rest := sum - longest
	if longest > rest+1 {
		return int64(2*rest + 1)
	}
	//otherwise , can complete all work
	return int64(sum)
}

//func solve(m []int, reserved int) int64 {
//	n := len(m)
//	sort.Ints(m)
//	week := 0
//
//	//cur := m[n-1]
//	cur := 0
//	for i := n - 1; i >= 0; i-- {
//		if cur == 0 {
//			cur = m[i]
//			continue
//		}
//		p := m[i]
//		if cur < p {
//			cur, p = p, cur
//		}
//		cur = cur - p
//		week += p + p
//	}
//
//	if cur > 0 {
//		if cur > reserved {
//			week += reserved + 1
//		} else {
//			week += cur
//		}
//	}
//	return int64(week)
//}
//
//func numberOfWeeks(m []int) int64 {
//	// 5,5, 7,7  的这种任务是个宝藏，可以充分利用。
//	cnts := make(map[int]int, len(m))
//	for _, x := range m {
//		cnts[x]++
//	}
//	reserved := 0
//	weeks := 0
//	for k, v := range cnts {
//		if v >= 2 {
//			reserved += k * (v / 2)
//		}
//		cnts[k] = v % 2
//	}
//
//	nums := make([]int, 0, len(m))
//	for k, v := range cnts {
//		if v != 0 {
//			nums = append(nums, k)
//		}
//	}
//
//	return solve(nums, reserved) + int64(reserved*2)
//}