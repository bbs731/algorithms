package simulation

// 这个解，太牛了！
func largestNumber(nums []int) string {
	ss := make([]string, len(nums))
	for i := range nums {
		ss[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i]+ss[j] >= ss[j]+ss[i]
	})
	if ss[0] == "0" {
		return "0"
	}

	return strings.Join(ss, "")
}

func largestNumber(nums []int) string {

	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})
	if ss[0] == "0" {
		return "0"
	}

	var ans strings.Builder
	for i := 0; i < len(nums); i++ {
		ans.WriteString(strconv.Itoa(nums[i]))
	}
	return ans.String()
}
