package problem0839

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Expected int
}

var Results = []Result{
	{[]string{"tars", "rats", "arts", "star"}, 2},
	{[]string{"omv", "ovm"}, 1},
	{[]string{"ABCDEF", "BACDEF", "CABDEF", "ACBDEF", "BCADEF", "CBDAEF", "BCDAEF", "DCBAEF", "CDBAEF", "BDCAEF", "DBCAEF", "DACBEF", "ADCBEF", "CDABEF", "DCABEF", "ACDBEF", "CADBEF", "BADCEF", "DABCEF", "BEACDF", "EBACDF", "ABECDF", "BAECDF", "BACEDF", "CBAEDF", "BCAEDF", "ACBEDF", "CABEDF", "CEBADF", "ECBADF", "BCEADF", "ACEBDF", "ECABDF", "CEABDF", "DEABCF", "EDABCF", "CEDFAB", "FEDCAB", "EFDCAB", "ECFADB", "CEFADB", "CAFEDB", "ACFEDB"}, 2},
}

func TestSimilarWordGroups(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := numSimilarGroups(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkSimilarWordGroups(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			numSimilarGroups(res.Input)
		}
	}
}

func TestSimilarWordGroupsOpt(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := numSimilarGroupsOpt(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkSimilarWordGroupsOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			numSimilarGroupsOpt(res.Input)
		}
	}
}
