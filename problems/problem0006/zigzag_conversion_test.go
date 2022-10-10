package problem0006

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Str      string
	NumRows  int
	Expected string
}

var Results = []Result{
	{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
	{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
	{"A", 3, "A"},
	{"A", 1, "A"},
	{"PAYPALISHIRING", 1, "PAYPALISHIRING"},
	{"PAYPALISHIRING", 2, "PYAIHRNAPLSIIG"},
	{"PAYPALISHIRING", 100, "PAYPALISHIRING"},
}

func TestZigZagConvert(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := convert(res.Str, res.NumRows)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
