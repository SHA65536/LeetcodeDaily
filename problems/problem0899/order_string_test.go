package problem0899

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	K        int
	Expected string
}

var Results = []Result{
	{"cba", 1, "acb"},
	{"baaca", 3, "aaabc"},
}

func TestWordSearchII(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := orderlyQueue(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
