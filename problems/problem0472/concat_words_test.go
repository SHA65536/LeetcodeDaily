package problem0472

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Expected []string
}

var Results = []Result{
	{[]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}, []string{"catsdogcats", "dogcatsdog", "ratcatdogcat"}},
	{[]string{"cat", "dog", "catdog"}, []string{"catdog"}},
}

func TestCheapestFlight(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findAllConcatenatedWordsInADict(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
