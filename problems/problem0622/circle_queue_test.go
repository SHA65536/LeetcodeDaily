package problem0622

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularQueue(t *testing.T) {
	assert := assert.New(t)

	q := Constructor(3)

	assert.Equal(true, q.EnQueue(1))
	assert.Equal(true, q.EnQueue(2))
	assert.Equal(true, q.EnQueue(3))
	assert.Equal(false, q.EnQueue(4))
	assert.Equal(3, q.Rear())
	assert.Equal(true, q.IsFull())
	assert.Equal(true, q.DeQueue())
	assert.Equal(true, q.EnQueue(4))
	assert.Equal(4, q.Rear())
	assert.Equal(true, q.DeQueue())
	assert.Equal(true, q.EnQueue(4))
	assert.Equal(4, q.Rear())
}
