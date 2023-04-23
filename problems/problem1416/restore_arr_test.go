package problem1416

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Limit    int
	Expected int
}

var Results = []Result{
	{"1000", 10000, 1},
	{"1000", 10, 0},
	{"1317", 2000, 8},
	{"1111111", 200, 44},
	{"1111111111111111111111111111", 2000, 54114452},
	{"600342244431311113256628376226052681059918526204", 703, 411743991},
}

func TestNumberOfPossibleArrays(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := numberOfArrays(res.Input, res.Limit)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
