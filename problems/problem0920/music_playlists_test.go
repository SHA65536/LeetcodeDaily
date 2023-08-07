package problem0920

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	N, G, K  int
	Expected int
}

var TestCases = []TestCase{
	{3, 3, 1, 6},
	{2, 3, 0, 6},
	{2, 3, 1, 2},
}

func TestNumberOfPlaylists(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := numMusicPlaylists(tc.N, tc.G, tc.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkNumberOfPlaylists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			numMusicPlaylists(tc.N, tc.G, tc.K)
		}
	}
}
