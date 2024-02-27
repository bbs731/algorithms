package interview_questions

// position hashing

func firstMissingPositive(nums []int) int {
	// Implement your solution here
	A := nums
	n := len(A)

	for i, num := range A {
		if num <= 0 || num >= n {
			continue
		}
		// swap
		for num > 0 && num <= n && num != A[num-1] {
			A[num-1], A[i] = A[i], A[num-1]
			num = A[i]
		}
	}

	for i := 1; i <= n; i++ {
		if A[i-1] != i {
			return i
		}
	}
	return n + 1
}
