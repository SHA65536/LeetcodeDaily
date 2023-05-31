package problem1396

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Commands []string
	Values   [][]any
	Expected []float64
}

var Results = []Result{
	{
		Commands: []string{"checkIn", "checkIn", "checkIn", "checkOut", "checkOut", "checkOut", "getAverageTime", "getAverageTime", "checkIn", "getAverageTime", "checkOut", "getAverageTime"},
		Values:   [][]any{{45, "Leyton", 3}, {32, "Paradise", 8}, {27, "Leyton", 10}, {45, "Waterloo", 15}, {27, "Waterloo", 20}, {32, "Cambridge", 22}, {"Paradise", "Cambridge"}, {"Leyton", "Waterloo"}, {10, "Leyton", 24}, {"Leyton", "Waterloo"}, {10, "Waterloo", 38}, {"Leyton", "Waterloo"}},
		Expected: []float64{0, 0, 0, 0, 0, 0, 14, 11, 0, 11, 0, 12},
	},
	{
		Commands: []string{"checkIn", "checkOut", "getAverageTime", "checkIn", "checkOut", "getAverageTime", "checkIn", "checkOut", "getAverageTime"},
		Values:   [][]any{{10, "Leyton", 3}, {10, "Paradise", 8}, {"Leyton", "Paradise"}, {5, "Leyton", 10}, {5, "Paradise", 16}, {"Leyton", "Paradise"}, {2, "Leyton", 21}, {2, "Paradise", 30}, {"Leyton", "Paradise"}},
		Expected: []float64{0, 0, 5.00000, 0, 0, 5.50000, 0, 0, 6.66667},
	},
}

func TestRailwaySystem(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := make([]float64, len(want))
		rail := Constructor()
		for i, cmd := range res.Commands {
			if cmd == "checkIn" {
				rail.CheckIn(res.Values[i][0].(int), res.Values[i][1].(string), res.Values[i][2].(int))
			} else if cmd == "checkOut" {
				rail.CheckOut(res.Values[i][0].(int), res.Values[i][1].(string), res.Values[i][2].(int))
			} else if cmd == "getAverageTime" {
				got[i] = rail.GetAverageTime(res.Values[i][0].(string), res.Values[i][1].(string))
			}
		}
		assert.InDeltaSlice(want, got, 0.0001, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkRailSystem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			rail := Constructor()
			for i, cmd := range res.Commands {
				if cmd == "checkIn" {
					rail.CheckIn(res.Values[i][0].(int), res.Values[i][1].(string), res.Values[i][2].(int))
				} else if cmd == "checkOut" {
					rail.CheckOut(res.Values[i][0].(int), res.Values[i][1].(string), res.Values[i][2].(int))
				} else if cmd == "getAverageTime" {
					rail.GetAverageTime(res.Values[i][0].(string), res.Values[i][1].(string))
				}
			}
		}
	}
}
