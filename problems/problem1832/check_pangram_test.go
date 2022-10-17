package problem1832

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected bool
}

var Results = []Result{
	{"thequickbrownfoxjumpsoverthelazydogaaaaaaaaaaa", true},
	{"thequickbrownfoxjumpsoverthelazydog", true},
	{"leetcode", false},
	{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false},
}

func TestCheckPangram(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := checkIfPangram(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
