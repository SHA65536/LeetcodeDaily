package problem0043

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	A, B     string
	Expected string
	Added    string
}

var Results = []Result{
	{"2", "3", "6", "5"},
	{"123", "456", "56088", "579"},
	{"123456", "654321", "80779853376", "777777"},
}

func TestMultiplyStrings(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := multiply(res.A, res.B)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func TestAddStrings(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		tA, tB, tE := []byte(res.A), []byte(res.B), []byte(res.Added)
		reverse(tA)
		reverse(tB)
		reverse(tE)
		got := add(tA, tB)
		assert.Equal(tE, got, fmt.Sprintf("%+v", res))
	}
}
