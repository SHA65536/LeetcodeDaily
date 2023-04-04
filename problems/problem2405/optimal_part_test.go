package problem2405

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected int
}

var Results = []Result{
	{"abacaba", 4},
	{"ssssss", 6},
	{"jvmzlipaqmqumaynpoaillexqthgmgsvdoyvvctcrppxcrgnqrmnfeoxlkqelshj", 12},
	{"gtbxtruaawbapxtfhgigrzocelaohxawqhvlgpfgomgesvjjslkzlmtzdmznrjve", 12},
	{"mxmfmtnvcqccldvddmfaetnttfcdjckdqowrpzacqwrheocsqaiayhkcngtgyjlj", 15},
	{"bxlwuhedwzhqdddlcilbpvauvsgjoydfjhoaldczrrdjhfszkkupwqbrciklflmp", 11},
	{"nfyxygbptvimpkckzoerkkkfbcimsltcoepqluchjwctmgusfcmkiticdsorjwul", 12},
	{"rojianbsiftkwwnpoeqanvxgebeisohyjuwgurcnuwuqlbypboxdcxvjdezczxpq", 11},
	{"bucqebmlygdnilutrtxvickudjgslgxmqhenlcwrlaurgbqiudbnkdmaahqybygp", 10},
	{"mtpclvjwfmbddfhzgikluwrixklqtcppmbahjlmbvlctqdpfpwgbmkvhvkvzvhjq", 10},
}

func TestPartitionStringMap(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := partitionStringMap(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkPartitionStringMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			partitionStringMap(res.Input)
		}
	}
}

func TestPartitionStringArr(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := partitionStringArr(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkPartitionStringArr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			partitionStringArr(res.Input)
		}
	}
}

func TestPartitionStringBit(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := partitionStringBit(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkPartitionStringBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			partitionStringBit(res.Input)
		}
	}
}
