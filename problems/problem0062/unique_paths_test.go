package problem0062

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	M, N     int
	Expected int
}

var Results = []Result{
	{3, 7, 28},
	{7, 3, 28},
	{3, 2, 3},
	{2, 3, 3},
	{100, 100, 4631081169483718960},
}

func TestUniquePathsBigInt(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := uniquePathsBigInt(res.M, res.N)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkUniquePathsBigInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			uniquePathsBigInt(res.M, res.N)
		}
	}
}
