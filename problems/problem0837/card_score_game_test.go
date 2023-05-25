package problem0837

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const MaxError float64 = 0.00001

type Result struct {
	N, K, Max int
	Expected  float64
}

var Results = []Result{
	{10, 1, 10, 1},
	{6, 1, 10, 0.6},
	{21, 17, 10, 0.73278},
}

func TestCardScoreGame(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := new21Game(res.N, res.K, res.Max)
		assert.LessOrEqual(math.Abs(got-want), MaxError, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkCardScoreGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			new21Game(res.N, res.K, res.Max)
		}
	}
}
