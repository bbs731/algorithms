package dp

import "sort"

// 周赛 124
/***
灵神的视频讲解：
https://www.bilibili.com/video/BV1Sm411U7cR/?spm_id_from=333.337.search-card.all.click&vd_source=84c3c489cf545fafdbeb3b3a6cd6a112
 */

// 简单的真是变态了， 凸显自己就是个猪脑子啊
func maxSelectedElements(nums []int) int {
	m := make(map[int]int) // 因为值域太大了， 所以不开数组，开map
	//n := len(nums)
	sort.Ints(nums)

	ans := 0
	for _, x := range nums {
		m[x+1] = m[x] + 1
		m[x] = m[x-1] + 1
		ans = max(ans, m[x+1], m[x])
	}
	return ans
}
