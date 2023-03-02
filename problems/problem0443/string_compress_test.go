package problem0443

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
	{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, []byte{'a', '2', 'b', '2', 'c', '3'}},
	{[]byte{'a'}, []byte{'a'}},
	{[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, []byte{'a', 'b', '1', '2'}},
}

func TestStringCompress(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		temp := make([]byte, len(res.Input))
		copy(temp, res.Input)
		got := compress(temp)
		assert.Equal(want, temp[:got], fmt.Sprintf("%+v", res))
	}
}
