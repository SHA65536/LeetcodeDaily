package problem1472

/*
You have a browser of one tab where you start on the homepage and you can visit another url,
get back in the history number of steps or move forward in the history number of steps.
Implement the BrowserHistory class:
    BrowserHistory(string homepage) Initializes the object with the homepage of the browser.
    void visit(string url) Visits url from the current page. It clears up all the forward history.
    string back(int steps) Move steps back in history. If you can only return x steps in the history and steps > x, you will return only x steps.
		Return the current url after moving back in history at most steps.
    string forward(int steps) Move steps forward in history. If you can only forward x steps in the history and steps > x,
		you will forward only x steps. Return the current url after forwarding in history at most steps.
*/

type BrowserHistory struct {
	History []string
	Pos     int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{History: []string{homepage}, Pos: 0}
}

func (b *BrowserHistory) Visit(url string) {
	b.History = b.History[:b.Pos+1]
	b.History = append(b.History, url)
	b.Pos++
}

func (b *BrowserHistory) Back(steps int) string {
	b.Pos = max(b.Pos-steps, 0)
	return b.History[b.Pos]
}

func (b *BrowserHistory) Forward(steps int) string {
	b.Pos = min(b.Pos+steps, len(b.History)-1)
	return b.History[b.Pos]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
