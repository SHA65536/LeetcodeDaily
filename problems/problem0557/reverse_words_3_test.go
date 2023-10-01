package problem0557

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Input    string
	Expected string
}

var Results = []Result{
	{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
	{"God Ding", "doG gniD"},
	{`According to all known lawsof aviation,there is no way a beeshould be able to fly.Its wings are too small to getits fat little body off the ground.The bee, of course, flies anywaybecause bees don't carewhat humans think is impossible.`,
		`gnidroccA ot lla nwonk foswal ereht,noitaiva si on yaw a dluohseeb eb elba ot stI.ylf sgniw era oot llams ot stiteg taf elttil ydob ffo eht ehT.dnuorg ,eeb fo ,esruoc seilf esuacebyawyna seeb t'nod tahwerac snamuh kniht si .elbissopmi`},
}

func TestReverseWords(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := reverseWords(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkReverseWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			reverseWords(res.Input)
		}
	}
}

func TestReverseWordsOpt(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		got := reverseWordsOpt(res.Input)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkReverseWordsOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			reverseWordsOpt(res.Input)
		}
	}
}
