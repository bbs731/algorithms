package bits_operation


// 用 &  and  ^  来实现  a + b  这是谁想到的？  巧妙啊
func getSum(a int, b int) int {
	for b != 0 {
		carry := (a&b) << 1
		a = a^b
		b = carry
	}

	return a
}
