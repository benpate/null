package null

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat(t *testing.T) {

	var f Float

	assert.True(t, f.IsNull())
	assert.False(t, f.IsPresent())
	assert.Zero(t, f.Float())
	assert.Equal(t, "", f.String())
	assert.Nil(t, f.Interface())

	// 1066 - Conquest of Anglo-Saxon England
	f.Set(1066.1014)
	assert.False(t, f.IsNull())
	assert.True(t, f.IsPresent())
	assert.Equal(t, 1066.1014, f.Float())
	assert.Equal(t, "1066.1014", f.String())
	assert.Equal(t, 1066.1014, f.Interface())

	// 1453 - Conquest of Contsantinople
	f.Set(1453.0402)
	assert.False(t, f.IsNull())
	assert.True(t, f.IsPresent())
	assert.Equal(t, 1453.0402, f.Float())
	assert.Equal(t, "1453.0402", f.String())
	assert.Equal(t, 1453.0402, f.Interface())

	f.Unset()
	assert.True(t, f.IsNull())
	assert.False(t, f.IsPresent())
	assert.Zero(t, f.Float())
	assert.Equal(t, "", f.String())
	assert.Nil(t, f.Interface())
}

func TestNewFloat(t *testing.T) {

	f := NewFloat(0)

	assert.False(t, f.IsNull())
	assert.True(t, f.IsPresent())
	assert.Zero(t, f.Float())
	assert.Equal(t, "0", f.String())

	// 1066 - Conquest of Anglo-Saxon England
	f.Set(1066.1014)
	assert.False(t, f.IsNull())
	assert.True(t, f.IsPresent())
	assert.Equal(t, 1066.1014, f.Float())
	assert.Equal(t, "1066.1014", f.String())

	// 1453 - Conquest of Contsantinople
	f.Set(1453.0402)
	assert.False(t, f.IsNull())
	assert.True(t, f.IsPresent())
	assert.Equal(t, 1453.0402, f.Float())
	assert.Equal(t, "1453.0402", f.String())

	f.Unset()
	assert.True(t, f.IsNull())
	assert.False(t, f.IsPresent())
	assert.Zero(t, f.Float())
	assert.Equal(t, "", f.String())
}
