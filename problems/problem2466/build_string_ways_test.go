package problem2466

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Low, High, Zero, One int
	Expected             int
}

var Results = []Result{
	{3, 3, 1, 1, 8},
	{2, 3, 1, 2, 5},
	{1, 1000, 1, 1, 376846411},
}

func TestCountWaysToBuildString(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := countGoodStrings(res.Low, res.High, res.Zero, res.One)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMaxQuizPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			countGoodStrings(res.Low, res.High, res.Zero, res.One)
		}
	}
}
