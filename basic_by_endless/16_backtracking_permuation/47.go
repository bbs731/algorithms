package backtracking

// hash function (这个东西有技巧吗， 去了解一下 hash function)
func convert(l []int) int {
	sum := 0
	for _, n := range l {
		sum = sum*30 + (n + 11)
	}
	return sum
}

func permuteUnique(nums []int) [][]int {
	ans := [][]int{}
	hans := make(map[int][]int)

	for _, a := range permute(nums) {
		hans[convert(a)] = a
	}

	for _, v := range hans {
		ans = append(ans, v)
	}
	return ans

}
