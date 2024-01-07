package weekly

func areaOfMaxDiagonal(dimensions [][]int) int {

	triangle := 0
	ans := 0
	for _, d := range dimensions {
		l, w := d[0], d[1]
		t := l*l + w*w
		if t > triangle {
			triangle = t
			ans = l*w
		}else if t == triangle {
			ans = max(ans, l*w)
		}
	}
	return ans
}
