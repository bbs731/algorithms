package bits_operation

import "math/bits"

/***
逻辑好简单啊，为啥想了半天呢？
 */
func minimizeXor(num1 int, num2 int) int {
	ones := bits.OnesCount(uint(num2))
	ans := 0
	for k := 30; k >= 0; k-- {
		if ones == 0 {
			break
		}
		if (1<<k)&num1 != 0 {
			// clear the bit (no need)
			//num1 &= ^(1 << k)
			ans |= 1 << k
			ones--
		}
	}
	if ones == 0 {
		return ans
	}

	// left ones, need to fill up
	for k := 0; ones > 0; k++ {
		if (1<<k)&ans == 0 {
			ones--
			ans |= 1 << k
		}
	}
	return ans
}

/***
链接：https://leetcode.cn/problems/minimize-xor/solutions/1864059/o1-kong-jian-fu-za-du-zuo-fa-by-endlessc-ywio/
灵神的答案，太牛了。思维的差距
 */

func minimizeXor(num1, num2 int) int {
	c1 := bits.OnesCount(uint(num1))
	c2 := bits.OnesCount(uint(num2))
	for ; c2 < c1; c2++ {
		num1 &= num1 - 1 // 最低的 1 变成 0
	}
	for ; c2 > c1; c2-- {
		num1 |= num1 + 1 // 最低的 0 变成 1
	}
	return num1
}
