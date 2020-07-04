package null

// Bool provides a nullable bool
type Bool struct {
	value   bool
	present bool
}

// NewBool returns a fully populated, nullable bool
func NewBool(value bool) Bool {
	return Bool{
		value:   value,
		present: true,
	}
}

// Bool returns the actual value of this object
func (b Bool) Bool() bool {
	return b.value
}

// Set applies a new value to the nullable item
func (b *Bool) Set(value bool) {
	b.value = value
	b.present = true
}

// Unset removes the value from this item, and sets it to null
func (b *Bool) Unset() {
	b.value = false
	b.present = false
}

// IsNull returns TRUE if this value is null
func (b Bool) IsNull() bool {
	return b.present == false
}

// IsPresent returns TRUE if this value is present
func (b Bool) IsPresent() bool {
	return b.present
}
