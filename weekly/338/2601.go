package weekly

func primeSubOperation(nums []int) bool {
	N := 1001
	primes := make([]bool, N)
	for i := 2; i < N; i++ {
		primes[i] = true
	}

	for i := 2; i < N; i++ {
		if primes[i] == false {
			continue
		}

		for j := 2 * i; j < N; j += i {
			primes[j] = false
		}
	}

	n := len(nums)
	last := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		found := false
		for j := last - 1; j > 0; j-- {
			if nums[i] <= j {
				last = nums[i]
				found = true
				break
			}

			if primes[nums[i]-j] == true {
				last = j
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
