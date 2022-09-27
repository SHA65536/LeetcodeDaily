package problem0917

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected string
}

var Results = []Result{
	{"ab-cd", "dc-ba"},
	{"a-bC-dEf-ghIj", "j-Ih-gfE-dCba"},
	{"Test1ng-Leet=code-Q!", "Qedo1ct-eeLg=ntse-T!"},
	{"-----", "-----"},
	{"abcde-", "edcba-"},
	{"a----", "a----"},
}

func TestReverseLettersOnly(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := reverseOnlyLetters(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
