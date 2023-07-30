package problem0808

/*
There are two types of soup: type A and type B.
Initially, we have n ml of each type of soup. There are four kinds of operations:
    Serve 100 ml of soup A and 0 ml of soup B,
    Serve 75 ml of soup A and 25 ml of soup B,
    Serve 50 ml of soup A and 50 ml of soup B, and
    Serve 25 ml of soup A and 75 ml of soup B.
When we serve some soup, we give it to someone, and we no longer have it.
Each turn, we will choose from the four operations with an equal probability 0.25.
If the remaining volume of soup is not enough to complete the operation, we will serve as much as possible.
We stop once we no longer have some quantity of both types of soup.
Note that we do not have an operation where all 100 ml's of soup B are used first.
Return the probability that soup A will be empty first, plus half the probability that A and B become empty at the same time.
Answers within 10-5 of the actual answer will be accepted.
*/

var servings = [][2]int{
	{100, 0}, {75, 25}, {50, 50}, {25, 75},
}

func soupServings(n int) float64 {
	if n >= 5000 {
		return 1
	}
	return solve(n, n, make(map[[2]int]float64))
}

func solve(a, b int, dp map[[2]int]float64) float64 {
	var res float64
	// Terminal states
	switch {
	case a <= 0 && b > 0: // More A
		return 1
	case a <= 0 && b <= 0: // Same amount
		return .5
	case a > 0 && b <= 0: // More B
		return 0
	}

	key := [2]int{a, b}
	if v, ok := dp[key]; ok {
		return v
	}

	for _, serve := range servings {
		res += solve(a-serve[0], b-serve[1], dp)
	}
	res *= .25
	dp[key] = res
	return res
}
