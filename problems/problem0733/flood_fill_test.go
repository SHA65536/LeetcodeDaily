package problem0733

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Image     [][]int
	Sr, Sc    int
	NewColour int
	Expected  [][]int
}

var Results = []Result{
	{
		Image: [][]int{{0, 0, 0}, {0, 0, 0}},
		Sr:    0, Sc: 0, NewColour: 2,
		Expected: [][]int{{2, 2, 2}, {2, 2, 2}},
	},
	{
		Image: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
		Sr:    1, Sc: 1, NewColour: 2,
		Expected: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
	},
}

func TestFloodFill(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := floodFill(res.Image, res.Sr, res.Sc, res.NewColour)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkFloodFill(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := make([][]int, 100)
		for i := 0; i < 100; i++ {
			input[i] = make([]int, 100)
		}
		floodFill(input, 0, 0, 2)
	}
}

/*
Constraints:
m == image.length
n == image[i].length
1 <= m, n <= 50
0 <= image[i][j], newColor < 216
0 <= sr < m
0 <= sc < n
*/
