package problem2353

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Cmds     []string
	Values   []Food
	Foods    []string
	Cuisines []string
	Ratings  []int
	Expected []string
}

var TestCases = []TestCase{
	{
		Cmds:     []string{"highestRated", "highestRated", "changeRating", "highestRated", "changeRating", "highestRated"},
		Foods:    []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"},
		Cuisines: []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"},
		Ratings:  []int{9, 12, 8, 15, 14, 7},
		Values:   []Food{{"korean", 0}, {"japanese", 0}, {"sushi", 16}, {"japanese", 0}, {"ramen", 16}, {"japanese", 0}},
		Expected: []string{"kimchi", "ramen", "", "sushi", "", "ramen"},
	},
}

func TestFoodRatings(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		var System = Constructor(tc.Foods, tc.Cuisines, tc.Ratings)
		var got = make([]string, len(tc.Cmds))
		var want = tc.Expected
		for i, cmd := range tc.Cmds {
			if cmd == "highestRated" {
				got[i] = System.HighestRated(tc.Values[i].Name)
			} else {
				System.ChangeRating(tc.Values[i].Name, tc.Values[i].Rating)
			}
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkSpecialPositionMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			var System = Constructor(tc.Foods, tc.Cuisines, tc.Ratings)
			for i, cmd := range tc.Cmds {
				if cmd == "highestRated" {
					System.HighestRated(tc.Values[i].Name)
				} else {
					System.ChangeRating(tc.Values[i].Name, tc.Values[i].Rating)
				}
			}
		}
	}
}
