package problem0609

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Expected [][]string
}

var Results = []Result{
	{[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"},
		[][]string{{"root/a/1.txt", "root/c/3.txt"}, {"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"}}},
	{[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)"},
		[][]string{{"root/a/1.txt", "root/c/3.txt"}, {"root/a/2.txt", "root/c/d/4.txt"}}},
}

func TestFileDuplicates(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := findDuplicate(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
