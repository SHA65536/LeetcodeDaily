package problem0072

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Source, Target string
	Expected       int
}

var Results = []Result{
	{"horse", "ros", 3},
	{"intention", "execution", 5},
}

func TestMinDeviationArray(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minDistance(res.Source, res.Target)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
