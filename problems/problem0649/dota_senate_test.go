package problem0649

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected string
}

var Results = []Result{
	{"RD", "Radiant"},
	{"RDD", "Dire"},
	{"DRRDDRRDRRRDRDRRDDRRRDRDDRDRRRDR", "Radiant"},
}

func TestArrayDifference(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := predictPartyVictory(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkArrayDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			predictPartyVictory(res.Input)
		}
	}
}
