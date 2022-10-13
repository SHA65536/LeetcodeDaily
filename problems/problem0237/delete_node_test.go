package problem0237

import (
	"fmt"
	. "leetcodedaily/helpers/listnode"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestTrianglePerimeter(t *testing.T) {
	assert := assert.New(t)

	head := MakeListNode(4, 5, 1, 9)
	res := MakeListNode(4, 1, 9)
	deleteNode(head.Next)
	assert.Equal(res, head, fmt.Sprintf("g:%s\ne:%s", head.String(), res.String()))

	head = MakeListNode(4, 5, 1, 9)
	res = MakeListNode(4, 5, 9)
	deleteNode(head.Next.Next)
	assert.Equal(res, head, fmt.Sprintf("g:%s\ne:%s", head.String(), res.String()))

	head = MakeListNode(4, 5, 1, 9)
	res = MakeListNode(5, 1, 9)
	deleteNode(head)
	assert.Equal(res, head, fmt.Sprintf("g:%s\ne:%s", head.String(), res.String()))
}
