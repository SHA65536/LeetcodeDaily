package problem1837

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int
	K        int
	Expected int
}

var Results = []Result{
	{34, 6, 9},
	{10, 10, 1},
	{1578, 7, 12},
	{98587423, 7, 37},
}

func TestSumOfDigitsInBaseK(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := sumBase(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkSumOfDigitsInBaseK(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			sumBase(res.Input, res.K)
		}
	}
}

func TestSumOfDigitsInBaseKMod(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := sumBaseMod(res.Input, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkSumOfDigitsInBaseKMod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			sumBaseMod(res.Input, res.K)
		}
	}
}
