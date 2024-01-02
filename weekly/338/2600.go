package weekly

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	if k <= numOnes {
		return k
	}
	if k <= numOnes+ numZeros {
		return numOnes
	}

	return 2*numOnes + numZeros - k
}
