package null

// Float provides a nullable float64
type Float struct {
	value   float64
	present bool
}

// NewFloat returns a fully populated, nullable float64
func NewFloat(value float64) Float {
	return Float{
		value:   value,
		present: true,
	}
}

// Get returns the actual value of this object
func (f Float) Get() float64 {
	return f.value
}

// Set applies a new value to the nullable item
func (f *Float) Set(value float64) {
	f.value = value
	f.present = true
}

// Unset removes the value from this item, and sets it to null
func (f *Float) Unset() {
	f.value = 0
	f.present = false
}

// IsNull returns TRUE if this value is null
func (f Float) IsNull() bool {
	return f.present == false
}

// IsPresent returns TRUE if this value is present
func (f Float) IsPresent() bool {
	return f.present
}
