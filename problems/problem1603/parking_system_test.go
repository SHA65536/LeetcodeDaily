package problem1603

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	B, M, S  int
	Cars     []int
	Expected []bool
}

var Results = []Result{
	{1, 1, 0, []int{1, 2, 3, 1}, []bool{true, true, false, false}},
}

func TestParkingSystem(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := make([]bool, len(res.Cars))
		parking := Constructor(res.B, res.M, res.S)
		for i := range res.Cars {
			got[i] = parking.AddCar(res.Cars[i])
		}
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkParkingSystem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			parking := Constructor(res.B, res.M, res.S)
			for i := range res.Cars {
				parking.AddCar(res.Cars[i])
			}
		}
	}
}
