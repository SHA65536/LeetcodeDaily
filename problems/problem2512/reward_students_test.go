package problem2512

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Result struct {
	Positive, Negative, Report []string
	StudentId                  []int
	K                          int
	Expected                   []int
}

var Results = []Result{
	{
		[]string{"smart", "brilliant", "studious"}, []string{"not"},
		[]string{"this student is studious", "the student is smart"},
		[]int{1, 2}, 2,
		[]int{1, 2},
	},
	{
		[]string{"smart", "brilliant", "studious"}, []string{"not"},
		[]string{"this student is not studious", "the student is smart"},
		[]int{1, 2}, 2,
		[]int{2, 1},
	},
	{
		[]string{"fkeofjpc", "qq", "iio"}, []string{"jdh", "khj", "eget", "rjstbhe", "yzyoatfyx", "wlinrrgcm"},
		[]string{"rjstbhe eget kctxcoub urrmkhlmi yniqafy fkeofjpc iio yzyoatfyx khj iio", "gpnhgabl qq qq fkeofjpc dflidshdb qq iio khj qq yzyoatfyx", "tizpzhlbyb eget z rjstbhe iio jdh jdh iptxh qq rjstbhe", "jtlghe wlinrrgcm jnkdbd k iio et rjstbhe iio qq jdh", "yp fkeofjpc lkhypcebox rjstbhe ewwykishv egzhne jdh y qq qq", "fu ql iio fkeofjpc jdh luspuy yzyoatfyx li qq v", "wlinrrgcm iio qq omnc sgkt tzgev iio iio qq qq", "d vhg qlj khj wlinrrgcm qq f jp zsmhkjokmb rjstbhe"},
		[]int{96537918, 589204657, 765963609, 613766496, 43871615, 189209587, 239084671, 908938263}, 3,
		[]int{239084671, 589204657, 43871615},
	},
}

func TestRewardStudentsMap(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		var studentcopy = make([]int, len(res.StudentId))
		copy(studentcopy, res.StudentId)
		got := topStudentsMap(res.Positive, res.Negative, res.Report, studentcopy, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkRewardStudentsMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			var studentcopy = make([]int, len(res.StudentId))
			copy(studentcopy, res.StudentId)
			topStudentsMap(res.Positive, res.Negative, res.Report, studentcopy, res.K)
		}
	}
}

func TestRewardStudentsTrie(t *testing.T) {
	assert := assert.New(t)

	for _, res := range Results {
		want := res.Expected
		var studentcopy = make([]int, len(res.StudentId))
		copy(studentcopy, res.StudentId)
		got := topStudentsTrie(res.Positive, res.Negative, res.Report, studentcopy, res.K)
		assert.Equal(want, got, fmt.Sprintf("%+v", res))
	}
}

func BenchmarkRewardStudentsTrie(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, res := range Results {
			var studentcopy = make([]int, len(res.StudentId))
			copy(studentcopy, res.StudentId)
			topStudentsTrie(res.Positive, res.Negative, res.Report, studentcopy, res.K)
		}
	}
}
