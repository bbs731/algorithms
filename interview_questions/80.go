package interview_questions

func removeDuplicates(nums []int) int {
	n := len(nums)
	start, end := 0, 0
	pos := 0
	for i:=1; i<n; i++{
		if nums[i] != nums[i-1]{
			num := nums[start]
			for k:=1; k<=min(2, end-start +1);k++ {
				nums[pos] = num
				pos++
			}
			start = i
			end = i
		} else {
			end++
		}
	}
	// last group of element [start, n-1]
	for k:=1; k<=min(2, n-start);k++ {
		nums[pos] = nums[n-1]
		pos++
	}
	return pos
}
