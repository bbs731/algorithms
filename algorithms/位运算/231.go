package bits_operation

import "math/bits"

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return bits.OnesCount(n) == 1
}
