package problem0319

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
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 1},
	{4, 2},
	{34, 5},
	{564896, 751},
}

func TestBulbSwitch(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := bulbSwitch(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkBulbSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			bulbSwitch(res.Input)
		}
	}
}

func TestBulbSwitchSqrt(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := bulbSwitchSqrt(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkBulbSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			bulbSwitchSqrt(res.Input)
		}
	}
}
