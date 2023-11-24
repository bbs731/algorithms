package backtracking

import (
	"strconv"
)

/*
isValidSolution 的代码，比 backtracking 的代码，还复杂
还容易出错。
 */

func isValidSolution(num, str1, str2 string) bool {
	n := len(num)
	pos := 0
	pos1, pos2 := len(str1), len(str2)
	n1, _ := strconv.Atoi(str1)
	n2, _ := strconv.Atoi(str2)
	sum := n1 + n2
	sums := strconv.Itoa(sum)
	suml := len(sums)

	//sanity check
	if pos1+pos2+suml > n {
		return false
	}

	if len(str1) == 0 || len(str2) == 0 {
		return false
	}

	if (str1[0] == '0' && len(str1) > 1) || (str2[0] == '0' && len(str2) > 1) {
		// no leading zero
		return false
	}

	for pos+pos1+pos2+suml <= n {
		if num[pos:pos+pos1+pos2+suml] == str1+str2+sums {
			pos = pos + pos1
			str1 = str2
			str2 = sums

			pos1, pos2 = len(str1), len(str2)
			n1, _ = strconv.Atoi(str1)
			n2, _ = strconv.Atoi(str2)
			sum = n1 + n2
			sums = strconv.Itoa(sum)
			suml = len(sums)
		} else {
			return false
		}
	}

	if pos+pos1+pos2 == n {
		return true
	}
	return false
}

func isAdditiveNumber(num string) bool {
	var dfs func(int, string, string) bool
	n := len(num)

	if n <= 2 {
		return false
	}

	dfs = func(i int, str1, str2 string) bool {
		if i == n {
			return isValidSolution(num, str1, str2)
		}

		// prune
		left := n - len(str1) - len(str2)
		if left < len(str1) || left < len(str2) {
			return false
		}

		if len(str2) == 0 {
			// select i to str1
			found := dfs(i+1, str1+string(num[i]), str2)
			if found {
				return true
			}
			// no select i to str1, add num[i] to str2
			if len(str1) > 0 {
				return dfs(i+1, str1, string(num[i]))
			}

		} else {
			// stop
			if isValidSolution(num, str1, str2) {
				return true
			}
			// no stop
			return dfs(i+1, str1, str2+string(num[i]))
		}
		return false
	}

	return dfs(0, "", "")
}
