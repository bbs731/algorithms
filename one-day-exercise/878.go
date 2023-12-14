package one_day_exercise

import "fmt"

/*
示例 1：

输入：n = 1, a = 2, b = 3
输出：2
示例 2：

输入：n = 4, a = 2, b = 3
输出：6
 */

/*
灵神的题解
https://leetcode.cn/problems/nth-magical-number/solutions/1984641/er-fen-da-an-rong-chi-yuan-li-by-endless-9j34/

这道题， gcd 写错一遍， 二分写错一遍!
 */
func nthMagicalNumber(n int, a int, b int) int {
	if a > b {
		a, b = b, a
	}

	c := gcd(a, b)
	lcm := a / c * b // minium multiplier   lcm = a *b /gcd(a,b)

	left := a
	right := n * a // [left, right]

	for left <= right {
		mid := left + (right-left)/2
		t := mid/a + mid/b - mid/lcm
		if t < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left % (int(1e9) + 7)
}

func gcd(a, b int) int {
	if b%a == 0 {
		return a
	}
	return gcd(b%a, a)
}

// gcd loop 版本
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

/*
思路是对的， 但是二分查找处理细节是错误的。
 */
func nthMagicalNumber(n int, a int, b int) int {
	if a > b {
		a, b = b, a
	}

	c := gcd(a, b)
	fmt.Printf("GCD: %d\n", c)
	mm := a * b / c // minium multiplier

	left := a
	right := n * a // [left, right]

	for left <= right {
		mid := left + (right-left)/2
		t := mid/a + mid/b - mid/mm
		if t > n {
			right = mid - 1
		} else if t < n {
			left = mid + 1
		} else {
			// 等于的情况， 其实 二分是可以自动找到 lower_bound 的。 不用这么麻烦。 看上面正确的代码。
			for true {
				if mid%a == 0 || mid%b == 0 {
					return mid % int(1e9+7)
				}
				mid--
			}
		}
	}
	return -1
}
