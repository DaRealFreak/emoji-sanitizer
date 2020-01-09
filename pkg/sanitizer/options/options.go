// Package options provides options for the unicode emoji sanitizer package
package options

// Option is the interface for all options
type Option interface {
	GetValue() interface{}
}

type failSilently bool

// GetValue implements the option interface method
func (o failSilently) GetValue() interface{} {
	return bool(o)
}

// FailSilently is the option to fail without quitting returning an exit code
func FailSilently(value bool) Option {
	return failSilently(value)
}

type loadFromOnline bool

// GetValue implements the option interface method
func (o loadFromOnline) GetValue() interface{} {
	return bool(o)
}

// LoadFromOnline is the option to download the emoji data from online instead of the local offline data
func LoadFromOnline(value bool) Option {
	return loadFromOnline(value)
}
