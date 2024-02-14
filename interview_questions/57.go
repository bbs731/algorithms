package interview_questions

/***
一出就死。
 */

//链接：https://leetcode.cn/problems/insert-interval/solutions/472435/shou-hua-tu-jie-57-cha-ru-qu-jian-fen-cheng

/***
写的真好， 不知道你的水平为啥会下降的这么厉害
 */

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	l := len(intervals)
	i := 0
	for i < l && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}
	for i < l && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	res = append(res, newInterval)
	for i < l {
		res = append(res, intervals[i])
		i++
	}
	return res
}

//func insert(intervals [][]int, newInterval []int) [][]int {
//	n := len(intervals)
//	ans := make([][]int, 0)
//	if n == 0 {
//		ans = append(ans, newInterval)
//		return ans
//	}
//	if newInterval[1] < intervals[0][0] {
//		ans = append([][]int{newInterval}, intervals...)
//		return ans
//	}
//	if newInterval[0] > intervals[n-1][1] {
//		ans = append(intervals, newInterval)
//		return ans
//	}
//	// use upper_bound
//	start := sort.Search(n, func(i int) bool {
//		//return intervals[i][0] >= newInterval[0]+1
//		return intervals[i][0] > newInterval[0]
//	})
//
//	if start == n || (start-1 >= 0 && intervals[start-1][1] > newInterval[0]) {
//		// start == 0 ? what happens?
//		start--
//	}
//
//	end := sort.Search(n, func(i int) bool {
//		return intervals[i][0] >= newInterval[1]
//	})
//
//	if end == n || (end-1 >= 0 && intervals[end-1][0] > newInterval[1]) {
//		end--
//	}
//
//	for i := 0; i < start; i++ {
//		ans = append(ans, intervals[i])
//	}
//	// merge interval
//	ans = append(ans, []int{min(intervals[start][0], newInterval[0]), max(intervals[end][1], newInterval[1])})
//
//	for i := end + 1; i < n; i++ {
//		ans = append(ans, intervals[i])
//	}
//	return ans
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
