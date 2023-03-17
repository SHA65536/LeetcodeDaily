package problem0208

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
		Commands: []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"},
		Values:   []string{"", "apple", "apple", "app", "app", "app", "app"},
		Expected: []bool{false, false, true, false, true, false, true},
	},
}

func TestMakeTrie(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := make([]bool, len(want))
		var trie Trie
		for i := range res.Commands {
			switch res.Commands[i] {
			case "Trie":
				trie = Constructor()
			case "insert":
				trie.Insert(res.Values[i])
			case "search":
				got[i] = trie.Search(res.Values[i])
			case "startsWith":
				got[i] = trie.StartsWith(res.Values[i])
			}
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
