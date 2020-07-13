package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testValue struct {
	Int   Int   `json:"int"`
	Float Float `json:"float"`
	Bool  Bool  `json:"bool"`
}

func TestUnmarshal_Empty(t *testing.T) {

	var value testValue

	j := []byte(`{}`)

	err := json.Unmarshal(j, &value)

	assert.Nil(t, err)
	assert.False(t, value.Int.present)
	assert.False(t, value.Float.present)
	assert.False(t, value.Bool.present)

	// Test zero values
	assert.Zero(t, value.Int.Int())
	assert.Zero(t, value.Float.Float())
	assert.False(t, value.Bool.Bool())
}

func TestUnmarshal_Nulls(t *testing.T) {

	var value testValue

	j := []byte(`{"int": null, "float": null, "bool":null}`)

	err := json.Unmarshal(j, &value)

	assert.Nil(t, err)
	assert.False(t, value.Int.present)
	assert.False(t, value.Float.present)
	assert.False(t, value.Bool.present)
}

func TestUnmarshal_Full(t *testing.T) {

	var value testValue

	j := []byte(`{"int": 1, "float": 3.14, "bool":true}`)

	err := json.Unmarshal(j, &value)

	assert.Nil(t, err)
	assert.True(t, value.Int.present)
	assert.Equal(t, 1, value.Int.Int())

	assert.True(t, value.Float.present)
	assert.Equal(t, 3.14, value.Float.Float())

	assert.True(t, value.Bool.present)
	assert.True(t, value.Bool.Bool())
}

func TestUnmarshal_FullAlt(t *testing.T) {

	var value testValue

	j := []byte(`{"float": 3, "bool": false}`)

	err := json.Unmarshal(j, &value)

	assert.Nil(t, err)
	assert.True(t, value.Float.present)
	assert.Equal(t, float64(3), value.Float.Float())

	assert.True(t, value.Bool.present)
	assert.False(t, value.Bool.Bool())
}

func TestUnmarshal_ErrorInt(t *testing.T) {

	var value testValue

	j := []byte(`{"int": "bad-value"}`)

	err := json.Unmarshal(j, &value)

	assert.NotNil(t, err)
	assert.False(t, value.Int.present)
}

func TestUnmarshal_ErrorFloat(t *testing.T) {

	var value testValue

	j := []byte(`{"float": "bad-value"}`)

	err := json.Unmarshal(j, &value)

	assert.NotNil(t, err)
	assert.False(t, value.Float.present)
}

func TestUnmarshal_ErrorBool(t *testing.T) {

	var value testValue

	j := []byte(`{"bool": "bad-value"}`)

	err := json.Unmarshal(j, &value)

	assert.NotNil(t, err)
	assert.False(t, value.Bool.present)
}

func TestMarshal_Empty(t *testing.T) {

	var value testValue

	result, err := json.Marshal(value)

	assert.Nil(t, err)
	assert.Equal(t, `{"int":null,"float":null,"bool":null}`, string(result))
}

func TestMarshal_Full(t *testing.T) {

	var value testValue

	value.Int.Set(1)
	value.Float.Set(3.14)
	value.Bool.Set(true)

	result, err := json.Marshal(value)

	assert.Nil(t, err)
	assert.Equal(t, `{"int":1,"float":3.14,"bool":true}`, string(result))
}
