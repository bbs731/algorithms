package bits_operation

func hasAlternatingBits(n int) bool {
	flip := 0
	if n&1 == 1 {
		flip = 1
	}
	for n > 0 {
		if n&1 != flip {
			return false
		}
		n = n >> 1
		if flip == 1 {
			flip = 0
		} else {
			flip = 1
		}
	}
	return true
}
