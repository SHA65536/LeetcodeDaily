package problem2284

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    []string
	Senders  []string
	Expected string
}

var Results = []Result{
	{[]string{"Hello userTwooo", "Hi userThree", "Wonderful day Alice", "Nice day userThree"}, []string{"Alice", "userTwo", "userThree", "Alice"}, "Alice"},
	{[]string{"How is leetcode for everyone", "Leetcode is useful for practice"}, []string{"Bob", "Charlie"}, "Charlie"},
	{[]string{"Wq Hb Iw HAq WhX PXj XC o ll", "ELw L mgN K qn yEP TmJ HGC TRK", "q kYO E dF l js q", "HV Mg Oic MT Rj yx CcO", "G Br d tTw S Rbt", "D", "bN y P LDX yq h Dg h", "caB J VDr du UD Ddp", "b DJI bR F Xjt", "m vyD Kfq qw", "aH", "i G SiF tk uzO LK JdE sO", "Fk lS H wW uq TmP TyY ZhN", "X", "iE bXc tf dZV", "G QYVc", "ydA VPV", "pei k zY dXi", "OAS rLN H h w Lhd aj lj RxR xf"}, []string{"yHfc", "ek", "aMatAkNUHm", "Bz", "RclslxNRo", "oD", "cRRZX", "HMzf", "RclslxNRo", "gUcyyFXiqc", "cRRZX", "ICUHL", "Bn", "OTBiDoT", "iE", "xPjhTIV", "gUcyyFXiqc", "Cmfa", "ofcOzZ"}, "RclslxNRo"},
}

func TestMostWordsSent(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := largestWordCount(res.Input, res.Senders)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}
