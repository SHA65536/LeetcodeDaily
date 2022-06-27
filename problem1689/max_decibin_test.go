package problem1689

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected int
}

var Results = []Result{
	{"12839739", 9},
	{"4", 4},
	{"123", 3},
	{"82734", 8},
	{"32", 3},
}

func TestMinPartitions(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := minPartitions(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
