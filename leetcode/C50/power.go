package C50

func recur_myPower(x float64, n int) float64 {
	if n == 1 {
		return x
	}

	res := recur_myPower(x, n/2)
	if n%2 == 0 {
		return res * res
	} else {
		return res * x * res
	}
}

func myPow(x float64, n int) float64 {
	var minus = false
	if n < 0 {
		minus = true
		n = -n
	} else if n == 0 {
		return 1
	}

	if minus {
		return 1 / recur_myPower(x, n)
	} else {
		return recur_myPower(x, n)
	}
}
