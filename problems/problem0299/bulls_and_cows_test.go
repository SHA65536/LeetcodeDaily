package problem0299

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Secret   string
	Guess    string
	Expected string
}

var Results = []Result{
	{"1807", "7810", "1A3B"},
	{"1123", "0111", "1A1B"},
	{
		"6938381421648648466378023698639070463545717322083780804052119010854594397751954185811429143960843601",
		"4472726215758441085421806561893114828949038540271489167947105010858115497347521909646492308453936736",
		"13A77B",
	},
}

func TestBullsAndCows(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := getHint(res.Secret, res.Guess)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkBullsAndCows(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			getHint(res.Secret, res.Guess)
		}
	}
}

func TestBullsAndCowsNoMap(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := getHintNoMap(res.Secret, res.Guess)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkBullsAndCowsNoMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			getHintNoMap(res.Secret, res.Guess)
		}
	}
}
