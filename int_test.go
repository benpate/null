package null

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {

	var b Int

	assert.True(t, b.IsNull())
	assert.False(t, b.IsPresent())
	assert.Zero(t, b.Get())

	// 753 BC - Founding of Rome
	b.Set(-753)
	assert.False(t, b.IsNull())
	assert.True(t, b.IsPresent())
	assert.Equal(t, -753, b.Get())

	// 410 AD - Fall of Rome
	b.Set(410)
	assert.False(t, b.IsNull())
	assert.True(t, b.IsPresent())
	assert.Equal(t, 410, b.Get())

	b.Unset()
	assert.True(t, b.IsNull())
	assert.False(t, b.IsPresent())
	assert.Zero(t, b.Get())

}

func TestNewInt(t *testing.T) {

	b := NewInt(0)
	assert.False(t, b.IsNull())
	assert.True(t, b.IsPresent())
	assert.Zero(t, b.Get())

	// 753 BC - Founding of Rome
	b.Set(-753)
	assert.False(t, b.IsNull())
	assert.True(t, b.IsPresent())
	assert.Equal(t, -753, b.Get())

	// 410 AD - Fall of Rome
	b.Set(410)
	assert.False(t, b.IsNull())
	assert.True(t, b.IsPresent())
	assert.Equal(t, 410, b.Get())

	b.Unset()
	assert.True(t, b.IsNull())
	assert.False(t, b.IsPresent())
	assert.Zero(t, b.Get())

}
