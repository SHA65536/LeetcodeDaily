package problem0211

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Commands []string
	Values   []string
	Expected []bool
}

var Results = []Result{
	{
		Commands: []string{"WordDictionary", "addWord", "addWord", "addWord", "search", "search", "search", "search"},
		Values:   []string{"", "bad", "dad", "mad", "pad", "bad", ".ad", "b.."},
		Expected: []bool{false, false, false, false, false, true, true, true},
	},
}

func TestMakeTrie(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := make([]bool, len(want))
		var wordDict WordDictionary
		for i := range res.Commands {
			switch res.Commands[i] {
			case "WordDictionary":
				wordDict = Constructor()
			case "addWord":
				wordDict.AddWord(res.Values[i])
			case "search":
				got[i] = wordDict.Search(res.Values[i])
			}
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
