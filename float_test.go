package null

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat(t *testing.T) {

	var b Float

	assert.True(t, b.IsNull())
	assert.False(t, b.IsPresent())
	assert.Zero(t, b.Get())

	// 1066 - Conquest of Anglo-Saxon England
	b.Set(1066.1014)
	assert.False(t, b.IsNull())
	assert.True(t, b.IsPresent())
	assert.Equal(t, 1066.1014, b.Get())

	// 1453 - Conquest of Contsantinople
	b.Set(1453.0402)
	assert.False(t, b.IsNull())
	assert.True(t, b.IsPresent())
	assert.Equal(t, 1453.0402, b.Get())

	b.Unset()
	assert.True(t, b.IsNull())
	assert.False(t, b.IsPresent())
	assert.Zero(t, b.Get())
}

func TestNewFloat(t *testing.T) {

	b := NewFloat(0)
	assert.False(t, b.IsNull())
	assert.True(t, b.IsPresent())
	assert.Zero(t, b.Get())

	// 1066 - Conquest of Anglo-Saxon England
	b.Set(1066.1014)
	assert.False(t, b.IsNull())
	assert.True(t, b.IsPresent())
	assert.Equal(t, 1066.1014, b.Get())

	// 1453 - Conquest of Contsantinople
	b.Set(1453.0402)
	assert.False(t, b.IsNull())
	assert.True(t, b.IsPresent())
	assert.Equal(t, 1453.0402, b.Get())

	b.Unset()
	assert.True(t, b.IsNull())
	assert.False(t, b.IsPresent())
	assert.Zero(t, b.Get())
}
