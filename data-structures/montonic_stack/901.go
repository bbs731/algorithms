package montonic_stack



type StockSpanner struct {
	s [][2]int // mono stack in decreasing order
}


func Constructor() StockSpanner {
	return StockSpanner{
		s: make([][2]int, 0),
	}
}


func (this *StockSpanner) Next(price int) int {
	ans := 1
	for len(this.s) > 0 && price >= (this.s[len(this.s)-1][0]) {
		ans += this.s[len(this.s)-1][1]
		// pop the last element
		this.s = this.s[:len(this.s)-1]
	}
	this.s = append(this.s, [2]int{price, ans})
	return ans
}

