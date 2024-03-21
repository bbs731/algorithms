package one_day_exercise

type FrequencyTracker struct {
	nums  map[int]int
	freqs map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{
		make(map[int]int),
		make(map[int]int),
	}

}

func (this *FrequencyTracker) Add(number int) {
	oldf := this.nums[number]
	this.nums[number]++
	newf := this.nums[number]

	if oldf != 0 {
		this.freqs[oldf]--
	}
	this.freqs[newf]++
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if _, ok := this.nums[number]; !ok {
		return
	}

	oldf := this.nums[number]
	this.freqs[oldf]--
	//  这段逻辑，特别容易被落下， debug 了好长时间
	if oldf != 1 {
		this.freqs[oldf-1]++
	}
	this.nums[number]--
	if this.nums[number] == 0 {
		delete(this.nums, number)
	}
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	//fmt.Println(this.nums)
	//fmt.Printf("freq:", this.freqs)
	return this.freqs[frequency] > 0
}
