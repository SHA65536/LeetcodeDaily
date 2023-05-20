package problem1357

/*
There is a supermarket that is frequented by many customers.
The products sold at the supermarket are represented as two parallel integer arrays products and prices,
where the ith product has an ID of products[i] and a price of prices[i].
When a customer is paying, their bill is represented as two parallel integer arrays product and amount,
where the jth product they purchased has an ID of product[j], and amount[j] is how much of the product they bought.
Their subtotal is calculated as the sum of each amount[j] * (price of the jth product).

The supermarket decided to have a sale.
Every nth customer paying for their groceries will be given a percentage discount.
The discount amount is given by discount, where they will be given discount percent off their subtotal.
More formally, if their subtotal is bill, then they would actually pay bill * ((100 - discount) / 100).
Implement the Cashier class:
    Cashier(int n, int discount, int[] products, int[] prices) Initializes the object with n,
	the discount, and the products and their prices.
    double getBill(int[] product, int[] amount) Returns the final total of the bill with the discount applied (if any).
	Answers within 10-5 of the actual value will be accepted.
*/

type Cashier struct {
	Every, Cur int
	Discount   float64
	Items      map[int]float64
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	var items = make(map[int]float64, len(products))
	for i := range products {
		items[products[i]] = float64(prices[i])
	}
	return Cashier{Every: n, Cur: 0, Discount: float64(discount), Items: items}
}

func (c *Cashier) GetBill(product []int, amount []int) float64 {
	var bill float64
	c.Cur++
	c.Cur %= c.Every
	for i := range product {
		bill += c.Items[product[i]] * float64(amount[i])
	}
	if c.Cur == 0 {
		bill *= (100 - c.Discount) / 100
	}
	return bill
}
