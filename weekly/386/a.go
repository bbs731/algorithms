package weekly

func isPossibleToSplit(nums []int) bool {
	cnts := make(map[int]int, len(nums))

	for _, n := range nums {
		cnts[n]++
		if cnts[n] >2 {
			return false
		}
	}
	return true
}
