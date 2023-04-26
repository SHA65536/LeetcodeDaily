package problem0258

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
	{38, 2},
	{0, 0},
	{2313548, 8},
	{9565614756, 9},
	{1245879654, 6},
}

func TestAddDigits(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := addDigits(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", got))
	}
}

func BenchmarkAddDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			addDigits(res.Input)
		}
	}
}

func TestAddDigitsMod(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := addDigitsMod(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", got))
	}
}

func BenchmarkAddDigitsMod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			addDigitsMod(res.Input)
		}
	}
}
