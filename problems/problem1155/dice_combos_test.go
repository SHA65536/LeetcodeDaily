package problem1155

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	N, K, T  int
	Expected int
}

var Results = []Result{
	{1, 6, 3, 1},
	{2, 6, 7, 6},
	{30, 30, 500, 222616187},
}

func TestDiceCombos(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := numRollsToTarget(res.N, res.K, res.T)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkDiceCombos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			numRollsToTarget(res.N, res.K, res.T)
		}
	}
}
