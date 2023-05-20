package problem0399

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Eqs      [][]string
	Vals     []float64
	Queries  [][]string
	Expected []float64
}

var Results = []Result{
	{
		[][]string{{"a", "b"}, {"b", "c"}},
		[]float64{2.0, 3.0},
		[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
		[]float64{6.0, 0.5, -1.0, 1.0, -1.0},
	},
	{
		[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
		[]float64{1.5, 2.5, 5.0},
		[][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
		[]float64{3.75000, 0.40000, 5.00000, 0.20000},
	},
	{
		[][]string{{"a", "b"}},
		[]float64{0.5},
		[][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
		[]float64{0.50000, 2.00000, -1.00000, -1.00000},
	},
}

func TestCalcDivEquation(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := calcEquation(res.Eqs, res.Vals, res.Queries)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkCalcDivEquation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			calcEquation(res.Eqs, res.Vals, res.Queries)
		}
	}
}
