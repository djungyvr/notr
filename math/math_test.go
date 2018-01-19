package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var maxInt32 int = 2147483647
var minInt32 int = -2147483648

var maxInt64 int64 = 9223372036854775807
var minInt64 int64 = -9223372036854775808

func TestMaxInt64(t *testing.T) {
	assert := assert.New(t)

	a, b := maxInt64, minInt64
	assert.Equal(MaxInt64(a, b), a, "max int64 is greater than min int64")

	a, b = maxInt64-1, maxInt64
	assert.Equal(MaxInt64(a, b), b, "max int64 is greater than max int64 - 1")

	a, b = -1, 1
	assert.Equal(MaxInt64(a, b), b, "1 is greater than -1")

	a, b = 1, 1
	assert.Equal(MaxInt64(a, b), b, "1 is equal to 1")

	a, b = 1, 1
	assert.Equal(MaxInt64(a, b), a, "1 is equal to 1")
}

func TestMaxInt(t *testing.T) {
	assert := assert.New(t)

	a, b := maxInt32, minInt32
	assert.Equal(MaxInt(a, b), a, "max int32 is greater than min int32")

	a, b = maxInt32-1, maxInt32
	assert.Equal(MaxInt(a, b), b, "max int32 is greater than max int32 - 1")

	a, b = -1, 1
	assert.Equal(MaxInt(a, b), b, "1 is greater than -1")

	a, b = 1, 1
	assert.Equal(MaxInt(a, b), b, "1 is equal to 1")

	a, b = 1, 1
	assert.Equal(MaxInt(a, b), a, "1 is equal to 1")
}

func TestMinInt(t *testing.T) {
	assert := assert.New(t)

	a, b := maxInt32, minInt32
	assert.Equal(MinInt(a, b), b, "min int32 is less than max int32")

	a, b = maxInt32-1, maxInt32
	assert.Equal(MinInt(a, b), a, "max int32 - 1 is less than max int32")

	a, b = -1, 1
	assert.Equal(MinInt(a, b), a, "-1 is less than 1")

	a, b = 1, 1
	assert.Equal(MinInt(a, b), b, "1 is equal to 1")

	a, b = 1, 1
	assert.Equal(MinInt(a, b), a, "1 is equal to 1")
}
