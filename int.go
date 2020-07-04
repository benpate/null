package null

// Int provides a nullable bool
type Int struct {
	value   int
	present bool
}

// NewInt returns a fully populated, nullable bool
func NewInt(value int) Int {
	return Int{
		value:   value,
		present: true,
	}
}

// Get returns the actual value of this object
func (i Int) Get() int {
	return i.value
}

// Set applies a new value to the nullable item
func (i *Int) Set(value int) {
	i.value = value
	i.present = true
}

// Unset removes the value from this item, and sets it to null
func (i *Int) Unset() {
	i.value = 0
	i.present = false
}

// IsNull returns TRUE if this value is null
func (i Int) IsNull() bool {
	return i.present == false
}

// IsPresent returns TRUE if this value is present
func (i Int) IsPresent() bool {
	return i.present
}
