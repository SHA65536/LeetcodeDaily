package problem0712

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	S1, S2   string
	Expected int
}

var TestCases = []TestCase{
	{"sea", "eat", 231},
	{"delete", "leet", 403},
	{"baheefggaichihabadheehbhhaabcdgfdeaiaadfgidbdidc", "agheibifadfibfcghefgeihafgefahdbhhbdgahhccefifdg", 5028},
	{"ihhhbhadabfeeigbfidaafhdicafhfebecgcbeciaefhggbd", "ddcehfhbdddfdcghaagaabbigabefabecdbaeeggaaddccfb", 5426},
	{"caifcdfeageaagcheceafgdcchggegiiiffcigcbafgiicib", "igeigceabigfbcbcaadaiffgfbeehigdceaahedahbfcbaba", 5468},
}

func TestMinAsciiDelete(t *testing.T) {
	assert := assert.New(t)

	for _, tc := range TestCases {
		want := tc.Expected
		got := minimumDeleteSum(tc.S1, tc.S2)
		assert.Equal(want, got, fmt.Sprintf("%+v", tc))
	}
}

func BenchmarkMinAsciiDelete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range TestCases {
			minimumDeleteSum(tc.S1, tc.S2)
		}
	}
}
