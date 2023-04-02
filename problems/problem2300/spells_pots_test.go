package problem2300

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Spells, Pots []int
	Success      int64
	Expected     []int
}

var Results = []Result{
	{[]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7, []int{4, 0, 3}},
	{[]int{3, 1, 2}, []int{8, 5, 8}, 16, []int{2, 0, 2}},
}

func TestSpellsAndPotions(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		tempS, tempP := make([]int, len(res.Spells)), make([]int, len(res.Pots))
		_, _ = copy(tempS, res.Spells), copy(tempP, res.Pots)
		got := successfulPairs(tempS, tempP, res.Success)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkSpellsAndPotions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			tempS, tempP := make([]int, len(res.Spells)), make([]int, len(res.Pots))
			successfulPairs(tempS, tempP, res.Success)
		}
	}
}

func TestSpellsAndPotionsTwoSum(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		tempS, tempP := make([]int, len(res.Spells)), make([]int, len(res.Pots))
		_, _ = copy(tempS, res.Spells), copy(tempP, res.Pots)
		got := successfulPairsTwoSum(tempS, tempP, res.Success)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkSpellsAndPotionsTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			tempS, tempP := make([]int, len(res.Spells)), make([]int, len(res.Pots))
			successfulPairsTwoSum(tempS, tempP, res.Success)
		}
	}
}
