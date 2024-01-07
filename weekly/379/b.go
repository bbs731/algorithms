package weekly


func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {

	if a == e  {
		if c != e || d < min(b,f) || d > max(b,f) {
			return 1
		}
	}
	if b == f {
		if b != d || c < min(a, e) || c > max(a, e) {
			return 1
		}
	}

	if c +d == e+f {
		if c+d != a+b {
			return 1
		}
		if a <min(c, e) || a > max(c, e){
			return 1
		}
	}

	if c-d == e-f {
		if c-d != a-b {
			return 1
		}

		if a < min(c, e) || a > max(c, e){
			return 1
		}
	}

	return 2

}
