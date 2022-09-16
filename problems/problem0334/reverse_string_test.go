package problem0334

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []byte
	Expected []byte
}

var Results = []Result{
	{[]byte("Hello"), []byte("olleH")},
	{[]byte("Helloo"), []byte("oolleH")},
	{[]byte("racecar"), []byte("racecar")},
	{[]byte("a"), []byte("a")},
	{[]byte(""), []byte("")},
}

func TestReverseString(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		reverseString(res.Input)
		assert.Equal(want, res.Input, fmt.Sprintf("%+v", res))
	}
}
