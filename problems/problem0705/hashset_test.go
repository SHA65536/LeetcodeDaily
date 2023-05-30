package problem0705

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Commands []string
	Values   []int
	Expected []bool
}

var Results = []Result{
	{
		Commands: []string{"add", "add", "contains", "contains", "add", "contains", "remove", "contains"},
		Values:   []int{1, 2, 1, 3, 2, 2, 2, 2},
		Expected: []bool{false, false, true, false, false, true, false, false},
	},
	{
		Commands: []string{"add", "add", "add", "contains", "contains", "contains", "remove", "contains", "contains", "contains", "remove", "contains", "contains", "contains"},
		Values:   []int{3, 203, 403, 3, 203, 403, 3, 3, 203, 403, 403, 3, 203, 403},
		Expected: []bool{false, false, false, true, true, true, false, false, true, true, false, false, true, false},
	},
}

func TestHashSet(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := make([]bool, len(want))
		set := Constructor()
		for i, cmd := range res.Commands {
			if cmd == "add" {
				set.Add(res.Values[i])
			} else if cmd == "remove" {
				set.Remove(res.Values[i])
			} else if cmd == "contains" {
				got[i] = set.Contains(res.Values[i])
			}
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkHashSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			set := Constructor()
			for i, cmd := range res.Commands {
				if cmd == "add" {
					set.Add(res.Values[i])
				} else if cmd == "remove" {
					set.Remove(res.Values[i])
				} else if cmd == "contains" {
					set.Contains(res.Values[i])
				}
			}
		}
	}
}
