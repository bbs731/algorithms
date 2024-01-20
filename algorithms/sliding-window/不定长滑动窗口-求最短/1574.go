package sliding_window

import (
	"fmt"
	"sort"
)

// 好题，经典的面试题目！

//https://leetcode.cn/problems/shortest-subarray-to-be-removed-to-make-array-sorted/solutions/2189149/dong-hua-yi-xie-jiu-cuo-liang-chong-xie-iijwz/
// 看看灵神的思路，是多他妈的清晰。

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	right := n - 1
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}
	if right == 0 { // arr 已经是非递减数组
		return 0
	}
	// 此时 arr[right-1] > arr[right]
	ans := right // 删除 arr[:right]
	for left := 0; ; right++ { // 枚举 right
		for right == n || arr[left] <= arr[right] {
			ans = min(ans, right-left-1) // 删除 arr[left+1:right]
			if arr[left] > arr[left+1] {
				return ans
			}
			left++
		}
	}
}

func min(a, b int) int { if a > b { return b }; return a }



// 这是非常好的面试题目啊， 第一次看到肯定过不去啊。
// tips 很有用啊。 不是最优的方法，还是复杂
func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	ascend := 1
	decend := n - 2


	// 得到的经验教训， 不要去两边都去考虑问题，这样问题的难度系数会翻倍！
	for ascend < n {
		if arr[ascend] < arr[ascend-1] {
			break
		}
		ascend++
	}
	ascend--
	// [0.. ascend] 升序

	for decend >= 0 {
		if arr[decend] > arr[decend+1] {
			break
		}
		decend--
	}
	decend++
	// [decend.. n-1] 降序

	fmt.Println(arr[:ascend+1])
	fmt.Println(arr[decend:])

	if ascend+1+n-decend > n {
		return 0
	}
	if arr[ascend] <= arr[decend] {
		return decend - 1 - ascend
	}

	//ans := n + 1   //test case [2, 2, 2, 1, 1, 1]  过不去
	ans := min(n - ascend -1, decend)

	// 找到 upper_bounds 最好  // 这个小聪明导致， 有些 test case 过不去。  [48, 88]   [64, 72]  这个应该去掉 88 但是，如果有 binary Search upperbound 你找到 88， 去掉 64, 72 答案是错误的。
	//l := sort.SearchInts(arr[:ascend+1], arr[decend])
	//l--
	//if l < 0 {
	//	l = 0
	//}
	l := 0
	r := decend

	for l <= ascend && r < n {
		for r < n && arr[r] < arr[l] {
			r++
		}
		ans = min(ans, r-l-1) // [l+1, r-1] 的部分需要删除
		l++
	}
	return ans
}
