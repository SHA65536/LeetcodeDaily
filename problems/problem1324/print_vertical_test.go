package problem1324

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected []string
}

var Results = []Result{
	{"HOW ARE YOU", []string{"HAY", "ORO", "WEU"}},
	{"TO BE OR NOT TO BE", []string{"TBONTB", "OEROOE", "   T"}},
	{"CONTEST IS COMING", []string{"CIC", "OSO", "N M", "T I", "E N", "S G", "T"}},
}

func TestPrintVertical(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := printVertically(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
