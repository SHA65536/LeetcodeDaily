package problem0049

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Expected [][]string
}

var Results = []Result{
	{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
	{[]string{""}, [][]string{{""}}},
	{[]string{"a"}, [][]string{{"a"}}},
	{[]string{"ddddddddddg", "dgggggggggg"}, [][]string{{"ddddddddddg"}, {"dgggggggggg"}}},
}

func TestGroupAnagrams(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := groupAnagrams(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkGroupAnagrams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			groupAnagrams(res.Input)
		}
	}
}

/*func TestGroupAnagramsOptInvalid(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := groupAnagramsOptInvalid(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkGroupAnagramsOptInvalid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			groupAnagramsOptInvalid(res.Input)
		}
	}
}*/

func TestGroupAnagramsOptValid(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := groupAnagramsOptValid(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkGroupAnagramsOptValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			groupAnagramsOptValid(res.Input)
		}
	}
}
