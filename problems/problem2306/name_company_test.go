package problem2306

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Expected int64
}

var Results = []Result{
	{[]string{"coffee", "donuts", "time", "toffee"}, 6},
	{[]string{"lack", "back"}, 0},
}

func TestNamingACompany(t *testing.T) {
	assert := assert.New(t)
	for _, res := range Results {
		want := res.Expected
		got := distinctNames(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
