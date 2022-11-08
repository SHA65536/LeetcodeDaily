package problem1408

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Expected []string
}

var Results = []Result{
	{[]string{"mass", "as", "hero", "superhero"}, []string{"as", "hero"}},
}

func TestMatchStrings(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := stringMatching(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
