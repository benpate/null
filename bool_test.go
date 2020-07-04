package null

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {

	var b Bool

	assert.True(t, b.IsNull())
	assert.False(t, b.IsPresent())
	assert.False(t, b.Get())

	b.Set(false)
	assert.False(t, b.IsNull())
	assert.True(t, b.IsPresent())
	assert.False(t, b.Get())

	b.Set(true)
	assert.False(t, b.IsNull())
	assert.True(t, b.IsPresent())
	assert.True(t, b.Get())

	b.Unset()
	assert.True(t, b.IsNull())
	assert.False(t, b.IsPresent())
	assert.False(t, b.Get())
}

func TestNewBool(t *testing.T) {

	b := NewBool(false)
	assert.False(t, b.IsNull())
	assert.True(t, b.IsPresent())
	assert.False(t, b.Get())

	b.Set(true)
	assert.False(t, b.IsNull())
	assert.True(t, b.IsPresent())
	assert.True(t, b.Get())

	b.Unset()
	assert.True(t, b.IsNull())
	assert.False(t, b.IsPresent())
	assert.False(t, b.Get())
}
