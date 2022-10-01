package problem0091

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected int
}

var Results = []Result{
	{"12", 2},
	{"226", 3},
	{"06", 0},
}

func TestSkylineProblem(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := numDecodings(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
