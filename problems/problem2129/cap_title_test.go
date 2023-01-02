package problem2129

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
	{"capiTalIze tHe titLe", "Capitalize The Title"},
	{"First leTTeR of EACH Word", "First Letter of Each Word"},
	{"i lOve leetcode", "i Love Leetcode"},
}

func TestCapitalizeTitle(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := capitalizeTitle(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkCapitalizeTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			capitalizeTitle(res.Input)
		}
	}
}

func TestCapitalizeTitleBit(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := capitalizeTitleBit(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkCapitalizeTitleBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			capitalizeTitleBit(res.Input)
		}
	}
}
