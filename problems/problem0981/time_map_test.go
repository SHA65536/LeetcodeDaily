package problem0981

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimestampMap(t *testing.T) {
	assert := assert.New(t)

	var timeMap = Constructor()

	timeMap.Set("foo", "bar", 1)
	assert.Equal("bar", timeMap.Get("foo", 1))
	assert.Equal("bar", timeMap.Get("foo", 3))
	timeMap.Set("foo", "bar2", 4)
	assert.Equal("bar2", timeMap.Get("foo", 4))
	assert.Equal("bar2", timeMap.Get("foo", 5))
	assert.Equal("", timeMap.Get("baz", 5))
}
