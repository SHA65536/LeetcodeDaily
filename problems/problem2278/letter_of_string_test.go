package problem2278

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Letter   byte
	Expected int
}

var Results = []Result{
	{"foobar", 'o', 33},
	{"jjjj", 'k', 0},
	{"jjjj", 'j', 100},
}

func TestLetterPercentInString(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := percentageLetter(res.Input, res.Letter)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
