package weekly

import "fmt"

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	n := len(bottomLeft)

	ans := 0
	for i:=0; i<n; i++ {
		for j:=i+1;j<n; j++ {
			a1, c1, b1, d1 := bottomLeft[i][0], topRight[i][0], bottomLeft[i][1], topRight[i][1]
			a2, c2, b2, d2 := bottomLeft[j][0], topRight[j][0], bottomLeft[j][1], topRight[j][1]

			if a2 >= c1 || c2 <=a1{
				// not intersect
				continue
			}
			if b2 >=d1 || b1 >=d2 {
				continue
			}

			var width, height int
			if a2 <= a1 && c2 >= c1 {
				// 包含关系
				width = c1-a1
			} else if a2 >=a1 && c2 <=c1 {
				width = c2-a2
			} else {
				width = min( abs(a2, c1), abs(a1, c2))
			}

			if b2 <= b1 && d2 >= d1 {
				height = d1-b1
			} else if b2 >=b1 && d2 <= d1 {
				height = d2-b2
			} else {
				height = min(abs(d1,b2), abs(b1, d2))
			}

			//fmt.Println(width, height)
			ans = max(ans, min(height, width))
		}
	}
	return int64(ans*ans)
}

func abs (a ,b int)int {
	if a >= b {
		return a -b
	}
	return b-a
}
