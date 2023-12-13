package monotonic_stack


type StockSpanner struct {
	s [][2]int // mono stack in decreasing order
}


func Constructor() StockSpanner {
	return StockSpanner{
		s: make([][2]int, 0),
	}
}


/*
比灵神给的答案简短。  但是灵神用了个哨兵的位置，可以不用比较 this.s 是空。
https://leetcode.cn/problems/online-stock-span/solutions/2470527/shi-pin-yi-ge-shi-pin-jiang-tou-dan-diao-cuk7/
 */
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


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
