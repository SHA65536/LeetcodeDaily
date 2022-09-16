package problem0458

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Buckets, MinutesDie, MinutesTest int
	Expected                         int
}

var Results = []Result{
	{1000, 15, 60, 5},
	{4, 15, 15, 2},
	{4, 15, 30, 2},
}

func TestPoorPigs(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := poorPigs(res.Buckets, res.MinutesDie, res.MinutesTest)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
