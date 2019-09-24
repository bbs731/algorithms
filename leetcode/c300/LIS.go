package c300

//Typical DP problem, Complexity is O(n^2) and have O(nlg(n)) algorithm

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 1
	l := make([]int, len(nums))

	for i, n := range nums {
		l[i] = 1
		for j := 0; j < i; j++ {
			if n >= nums[j] && l[j]+1 > l[i] {
				l[i] = l[j] + 1
				if l[i] > max {
					max = l[i]
				}
			}
		}
	}
	return max
}

/*
 * func LengthOfLIS(nums []int) (int, int, []int) {
 *     var index int
 *     max := 1
 *     p := make([]int, len(nums))
 *     l := make([]int, len(nums))
 *
 *     for i, _ := range nums {
 *         l[i] = 1
 *         for j := 0; j < i; j++ {
 *             if nums[i] >= nums[j] && l[j]+1 > l[i] {
 *                 l[i] = l[j] + 1
 *                 p[i] = j
 *                 if l[i] > max {
 *                     max = l[i]
 *                     index = i
 *                 }
 *             }
 *         }
 *     }
 *
 *     return max, index, p
 * }
 */
