package problem1323

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int
	Expected int
}

var Results = []Result{
	{9996, 9999},
	{9669, 9969},
	{9999, 9999},
}

func TestMaximumNumber(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maximum69Number(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMaximumNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			maximum69Number(res.Input)
		}
	}
}

func TestMaximumNumberOpt(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := maximum69NumberOpt(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkMaximumNumberOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			maximum69NumberOpt(res.Input)
		}
	}
}
