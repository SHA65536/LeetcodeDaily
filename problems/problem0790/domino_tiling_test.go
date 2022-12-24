package problem0790

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    int
	Expected int
}

var Results = []Result{
	{3, 5},
	{1, 1},
}

func TestTrominoDomino(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := numTilings(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
