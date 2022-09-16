package problem1220

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
	{1, 5},
	{2, 10},
	{5, 68},
	{10, 1739},
}

func TestCountVowelPermutation(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := countVowelPermutation(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkVowelPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			countVowelPermutation(res.Input)
		}
	}
}

func TestCountVowelPermutationNaive(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := countVowelPermutationNaive(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkVowelPermutationNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			countVowelPermutationNaive(res.Input)
		}
	}
}
