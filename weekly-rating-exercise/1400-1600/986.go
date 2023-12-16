package _400_1600

// 好题， 就是太简单。 锻炼变成，和边界条件的锻炼。 1542分
func intersect(a, b, c, d int) ([]int, bool) {
	if b < c || d < a {
		return nil, false
	}
	// 这里还能简化逻辑.  return max(a,c), min(b,d)
	//if a <= c {
	//	return []int{c, min(b, d)}, true
	//}
	//return []int{a, min(b, d)}, true

	return []int{max(a,c), min(b,d)}, true
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	ans := make([][]int, 0)
	l, r := 0, 0

	for l < len(firstList) && r < len(secondList) {
		if interval, ok := intersect(firstList[l][0], firstList[l][1], secondList[r][0], secondList[r][1]); ok {
			ans = append(ans, interval)
		}

		if firstList[l][1] < secondList[r][1] {  // 这里是个坑。容易错。 需用用 end 来移动, 而不是 start
			l++
		} else {
			r++
		}
	}

	return ans
}
