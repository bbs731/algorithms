package weekly

//You are given a 0-indexed 2D integer array nums representing the coordinates o
//f the cars parking on a number line. For any index i, nums[i] = [starti, endi] w
//here starti is the starting point of the ith car and endi is the ending point of
// the ith car.
//
// Return the number of integer points on the line that are covered with any par
//t of a car.
//
//
// Example 1:
//
//
//Input: nums = [[3,6],[1,5],[4,7]]
//Output: 7
//Explanation: All the points from 1 to 7 intersect at least one car, therefore
//the answer would be 7.
//
//
// Example 2:
//
//
//Input: nums = [[1,3],[5,8]]
//Output: 7
//Explanation: Points intersecting at least one car are 1, 2, 3, 5, 6, 7, 8. The
//re are a total of 7 points, therefore the answer would be 7.
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 100
// nums[i].length == 2
// 1 <= starti <= endi <= 100
//
// Related Topics Hash Table Math Prefix Sum
/*
https://leetcode.cn/problems/points-that-intersect-with-cars/solutions/2435384/chai-fen-shu-zu-xian-xing-zuo-fa-by-endl-3xpm/
思路太牛了， 赞啊灵神
 */
func numberOfPoints(nums [][]int) int {
	// 这里可以简化成 maxn 直接去 102
	//maxn := 0
	//for _, r := range nums {
	//	maxn = max(maxn, r[1])
	//}
	// 根据灵神的想法， 创建差分数组, 可以取名 diff
	p := make([]int, 100+2)

	for _, r := range nums {
		p[r[0]]++
		p[r[1]+1]--
	}

	ans := 0
	// 这里可以简化，因为所有线段范围都是最小从 1 开始的。
	//if p[0] >= 1 {
	//	ans++
	//}
	// 请前缀和等到原来的数组。 因为我们给每一个段 [a, b] 里面都加了 1， 所以只需要判断， 这个位置的数是不是 > 1， 就知道是不是这个位置被某个区间覆盖了。
	for i := 1; i <= 101; i++ {
		p[i] += p[i-1]
		if p[i] >= 1 {
			ans++
		}
	}
	return ans
}
