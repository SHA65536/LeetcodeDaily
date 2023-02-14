package problem0067

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	A, B     string
	Expected string
}

var Results = []Result{
	{"11", "1", "100"},
	{"1010", "1011", "10101"},
}

func TestAddBinaryStrings(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := addBinary(res.A, res.B)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
