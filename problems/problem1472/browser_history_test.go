package problem1472

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Commands []string
	Values   []any
	Expected []string
}

var Results = []Result{
	{
		[]string{"BrowserHistory", "visit", "visit", "visit", "back", "back", "forward", "visit", "forward", "back", "back"},
		[]any{"leetcode.com", "google.com", "facebook.com", "youtube.com", 1, 1, 1, "linkedin.com", 2, 2, 7},
		[]string{"", "", "", "", "facebook.com", "google.com", "facebook.com", "", "linkedin.com", "google.com", "leetcode.com"},
	},
}

func TestBrowserHistory(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := make([]string, len(want))
		var browser BrowserHistory
		for i := range res.Commands {
			switch res.Commands[i] {
			case "BrowserHistory":
				browser = Constructor(res.Values[i].(string))
			case "visit":
				browser.Visit(res.Values[i].(string))
			case "back":
				got[i] = browser.Back(res.Values[i].(int))
			case "forward":
				got[i] = browser.Forward(res.Values[i].(int))
			}
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
