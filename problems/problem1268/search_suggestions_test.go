package problem1268

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Search   string
	Expected [][]string
}

var Results = []Result{
	{
		[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
		"mouse",
		[][]string{
			{"mobile", "moneypot", "monitor"},
			{"mobile", "moneypot", "monitor"},
			{"mouse", "mousepad"},
			{"mouse", "mousepad"},
			{"mouse", "mousepad"},
		},
	},
	{
		[]string{"havvana"},
		"havvana",
		[][]string{
			{"havvana"},
			{"havvana"},
			{"havvana"},
			{"havvana"},
			{"havvana"},
			{"havvana"},
			{"havvana"},
		},
	},
	{
		[]string{"bags", "baggage", "banner", "box", "cloths"},
		"bags",
		[][]string{
			{"baggage", "bags", "banner"},
			{"baggage", "bags", "banner"},
			{"baggage", "bags"},
			{"bags"},
		},
	},
}

func TestSuggestedProducts(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := suggestedProducts(res.Input, res.Search)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
