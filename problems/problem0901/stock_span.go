package problem0901

/*
Design an algorithm that collects daily price quotes for some stock and returns the span of that stock's price for the current day.
The span of the stock's price today is defined as the maximum number of consecutive days
(starting from today and going backward) for which the stock price was less than or equal to today's price.
For example, if the price of a stock over the next 7 days were [100,80,60,70,60,75,85], then the stock spans would be [1,1,1,2,1,4,6].
Implement the StockSpanner class:
StockSpanner() Initializes the object of the class.
int next(int price) Returns the span of the stock's price given that today's price is price.
*/

type StockSpanner struct {
	Vals []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		Vals: []int{},
	}
}

func (this *StockSpanner) Next(price int) int {
	var count int
	this.Vals = append(this.Vals, price)
	for i := len(this.Vals) - 1; i >= 0; i-- {
		if this.Vals[i] > price {
			return count
		}
		count++
	}
	return count
}

type StockSpannerOpt struct {
	Vals [][2]int
}

func ConstructorOpt() StockSpannerOpt {
	return StockSpannerOpt{
		Vals: [][2]int{},
	}
}

func (this *StockSpannerOpt) NextOpt(price int) int {
	var res int = 1
	for len(this.Vals) > 0 && this.Vals[len(this.Vals)-1][0] <= price {
		res += this.Vals[len(this.Vals)-1][1]
		this.Vals = this.Vals[:len(this.Vals)-1]
	}
	this.Vals = append(this.Vals, [2]int{price, res})
	return res
}
