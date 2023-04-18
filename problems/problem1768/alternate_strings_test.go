package problem1768

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	InA, InB string
	Expected string
}

var Results = []Result{
	{"abc", "pqr", "apbqcr"},
	{"ab", "pqrs", "apbqrs"},
	{"abcd", "pq", "apbqcd"},
	{
		"wwhedtnmxnwvjzkorokrabwfgcxvhvvrbcjyovjjymtcaksykugkcbinkhiwnxzxcndvjxvfijsepvfkxdxewxukqcedfjtoeiro",
		"mjcrztcxbszwtryzgkllofvxmsarumawnpowlpqlwlpgtmmatnnhynlbrlsxsblcafheilmclghrfoitqwndqwavumhqtclctayb",
		"wmwjhcerdzttncmxxbnswzvwjtzrkyozrgokklrlaobfwvfxgmcsxavrhuvmvarwbncpjoywolvpjqjlywmltpcgatkmsmyaktungnkhcybnilnbkrhliswxnsxbzlxccanfdhvejixlvmfciljgsherpfvofiktxqdwxnedwqxwuakvqucmehdqftjctlocetiaryob",
	},
	{
		"zygzhjgprfrhoxbygzngjepdwlfendfxsztjtaccncmsmlnndddcgpbnrjebcqffasunfcflyaavfhzrldedkgrltptbkbffdaar",
		"vnxrpvrpxwfmklraeczeinqvqwrmkvqwrefqjjjeqxythilphhrsxjmkkhafgfcdnvvwvpwkjxaxybqphwuselyibmrozjqrrmiu",
		"zvyngxzrhpjvgrpprxfwrfhmokxlbryagezcnzgejienpqdvwqlwfremnkdvfqxwsrzetfjqtjajcjcenqcxmystmhlinlnpdhdhdrcsgxpjbmnkrkjheabfcgqffcfdansvuvnwfvcpfwlkyjaxaavxfyhbzqrplhdweudskeglrylitbpmtrbokzbjfqfrdramairu",
	},
}

func TestAlternateStringsBuilder(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := mergeAlternatelyBuilder(res.InA, res.InB)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkAlternateStringsBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			mergeAlternatelyBuilder(res.InA, res.InB)
		}
	}
}

func TestAlternateStringsSlice(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := mergeAlternatelySlice(res.InA, res.InB)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkAlternateStringsSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			mergeAlternatelySlice(res.InA, res.InB)
		}
	}
}

func TestAlternateStringsUnsafe(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := mergeAlternatelyUnsafe(res.InA, res.InB)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkAlternateStringsUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			mergeAlternatelyUnsafe(res.InA, res.InB)
		}
	}
}
