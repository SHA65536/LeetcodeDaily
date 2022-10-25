package problem1662

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input1   []string
	Input2   []string
	Expected bool
}

var Results = []Result{
	{[]string{"ab", "c"}, []string{"a", "bc"}, true},
	{[]string{"a", "cb"}, []string{"ab", "c"}, false},
	{[]string{"abc", "d", "defg"}, []string{"abcddefg"}, true},
	{[]string{"abc", "d", "defg"}, []string{"abcddefgijklmn"}, false},
	{[]string{"abc", "d", "defg", "hijkl"}, []string{"abcddefg"}, false},
}

func TestArrayStringsAreEqual(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := arrayStringsAreEqual(res.Input1, res.Input2)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkArrayStringsAreEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			arrayStringsAreEqual(res.Input1, res.Input2)
		}
	}
}

func TestArrayStringsAreEqualNaive(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := arrayStringsAreEqualNaive(res.Input1, res.Input2)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkArrayStringsAreEqualNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			arrayStringsAreEqualNaive(res.Input1, res.Input2)
		}
	}
}
