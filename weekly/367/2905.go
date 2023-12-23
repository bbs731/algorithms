package weekly

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	n := len(nums)
	maxIndex := 0
	minIndex := 0

	for j:=indexDifference; j<n; j++ {
		i := j - indexDifference
		if nums[i] > nums[maxIndex]{
			maxIndex =i
		}else if nums[i] < nums[minIndex]{
			minIndex = i
		}

		if abs(nums[j] - nums[minIndex]) >= valueDifference {
			return []int{minIndex, j}
		}
		if abs(nums[maxIndex] - nums[j]) >= valueDifference{
			return []int {maxIndex, j}
		}
	}
	return []int{-1, -1}
}

func abs(a int)int {
	if  a < 0 {
		return -a
	}
	return a
}