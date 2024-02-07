package bits_operation

import "math/bits"

// 好玩！  1 is power of 4
func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	return bits.OnesCount(uint(n)) == 1 && bits.TrailingZeros(uint(n))%2 == 0
}
