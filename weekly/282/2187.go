package weekly

import "sort"


func cost (time []int, m int)int {
	total :=0
	for i := range time{
		total += m/time[i]
	}
	return total
}

func minimumTime(time []int, totalTrips int) int64 {
	sort.Ints(time)

	l, r := 0, time[0]*totalTrips
	for l<=r {
		mid := (l+r)>>1
		total := cost(time, mid)
		if total < totalTrips {
			l = mid+1
		} else {
			r = mid-1
		}
	}
	return int64 (l)
}
