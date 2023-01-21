package problem0093

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
	{"25525511135", []string{"255.255.11.135", "255.255.111.35"}},
	{"0000", []string{"0.0.0.0"}},
	{"101023", []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}},
	{"0", []string{}},
}

func TestRestoreAddresses(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := restoreIpAddresses(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkRestoreAddresses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			restoreIpAddresses(res.Input)
		}
	}
}
