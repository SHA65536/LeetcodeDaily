package problem0858

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	P, Q     int
	Expected int
}

var Results = []Result{
	{2, 1, 2},
	{3, 1, 1},
}

func TestMirrorReflection(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := mirrorReflection(res.P, res.Q)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
