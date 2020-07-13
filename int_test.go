package null

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {

	var i Int

	assert.True(t, i.IsNull())
	assert.False(t, i.IsPresent())
	assert.Zero(t, i.Int())
	assert.Equal(t, "", i.String())

	// 753 BC - Founding of Rome
	i.Set(-753)
	assert.False(t, i.IsNull())
	assert.True(t, i.IsPresent())
	assert.Equal(t, -753, i.Int())
	assert.Equal(t, "-753", i.String())

	// 410 AD - Fall of Rome
	i.Set(410)
	assert.False(t, i.IsNull())
	assert.True(t, i.IsPresent())
	assert.Equal(t, 410, i.Int())
	assert.Equal(t, "410", i.String())

	i.Unset()
	assert.True(t, i.IsNull())
	assert.False(t, i.IsPresent())
	assert.Zero(t, i.Int())
	assert.Equal(t, "", i.String())
}

func TestNewInt(t *testing.T) {

	i := NewInt(0)

	assert.False(t, i.IsNull())
	assert.True(t, i.IsPresent())
	assert.Zero(t, i.Int())
	assert.Equal(t, "0", i.String())

	// 753 BC - Founding of Rome
	i.Set(-753)
	assert.False(t, i.IsNull())
	assert.True(t, i.IsPresent())
	assert.Equal(t, -753, i.Int())
	assert.Equal(t, "-753", i.String())

	// 410 AD - Fall of Rome
	i.Set(410)
	assert.False(t, i.IsNull())
	assert.True(t, i.IsPresent())
	assert.Equal(t, 410, i.Int())
	assert.Equal(t, "410", i.String())

	i.Unset()
	assert.True(t, i.IsNull())
	assert.False(t, i.IsPresent())
	assert.Zero(t, i.Int())
	assert.Equal(t, "", i.String())
}
