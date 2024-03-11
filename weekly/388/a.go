package  weekly

import "sort"

func minimumBoxes(apple []int, capacity []int) int {
	s := 0
	for _, a := range apple {
		s +=a
	}

	sort.Ints(capacity)
	cnts := 0
	for i:=len(capacity)-1; i>=0; i-- {
		if capacity[i] >= s{
			cnts++
			break
		} else {
			s -= capacity[i]
		}
	}
	return cnts
}