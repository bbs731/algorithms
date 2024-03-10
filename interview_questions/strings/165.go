package strings

import (
	"strconv"
	"strings"
)

/***
这个 ，写的，太好了。
 */

func compareVersion(version1 string, version2 string) int {
	n, m := len(version1), len(version2)
	i, j := 0, 0
	for i < n || j < m {
		x, y := 0, 0
		for ; i < n && version1[i] != '.'; i++ {
			x = x*10 + int(version1[i]-'0')
		}
		i++
		for ; j < m && version2[j] != '.'; j++ {
			y = y*10 + int(version2[j]-'0')
		}
		j++
		if x > y {
			return 1
		} else if x < y {
			return -1
		}
	}

	return 0
}


/***
这种逻辑， 就特别容易出错， corner case 考虑不到！
 */
func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	var i int
	for ; i<len(v1) && i<len(v2); {
		n1, _ := strconv.Atoi(v1[i])
		n2, _ := strconv.Atoi(v2[i])
		if n1 > n2 {
			return 1
		} else if n1 < n2 {
			return -1
		}
		i++
	}
	ans := 0
	if len(v2) > len(v1) {
		ans = -1
		v1 , v2 = v2, v1
	} else if len(v2) < len(v1) {
		ans = 1
	} else {
		return 0
	}

	for ;i <len(v1); i++ {
		n, _ := strconv.Atoi(v1[i])
		if n != 0 {
			return ans
		}
	}
	return 0
}
