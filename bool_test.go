package null

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {

	var b Bool

	assert.True(t, b.IsNull())
	assert.False(t, b.IsPresent())
	assert.False(t, b.Bool())
	assert.Equal(t, "", b.String())

	b.Set(false)
	assert.False(t, b.IsNull())
	assert.True(t, b.IsPresent())
	assert.False(t, b.Bool())
	assert.Equal(t, "false", b.String())

	b.Set(true)
	assert.False(t, b.IsNull())
	assert.True(t, b.IsPresent())
	assert.True(t, b.Bool())
	assert.Equal(t, "true", b.String())

	b.Unset()
	assert.True(t, b.IsNull())
	assert.False(t, b.IsPresent())
	assert.False(t, b.Bool())
	assert.Equal(t, "", b.String())
}

func TestNewBool(t *testing.T) {

	b := NewBool(false)
	assert.False(t, b.IsNull())
	assert.True(t, b.IsPresent())
	assert.False(t, b.Bool())

	b.Set(true)
	assert.False(t, b.IsNull())
	assert.True(t, b.IsPresent())
	assert.True(t, b.Bool())

	b.Unset()
	assert.True(t, b.IsNull())
	assert.False(t, b.IsPresent())
	assert.False(t, b.Bool())
}
