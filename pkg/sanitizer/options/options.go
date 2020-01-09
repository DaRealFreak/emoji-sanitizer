// Package options provides options for the unicode emoji sanitizer package
package options

// Option is the interface for all options
type Option interface {
	GetValue() interface{}
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

type loadCustomPath string

// GetValue implements the option interface method
func (o loadCustomPath) GetValue() interface{} {
	return string(o)
}

// LoadFromCustomPath is the option to download the emoji data from a custom path, can be either URL or file path
func LoadFromCustomPath(value string) Option {
	return loadCustomPath(value)
}

type allowEmojiCodes []string

// GetValue implements the option interface method
func (o allowEmojiCodes) GetValue() interface{} {
	return []string(o)
}

// AllowEmojiCodes is the option to allow specific emoji codes (like numbers, etc...)
func AllowEmojiCodes(codes []string) Option {
	return allowEmojiCodes(codes)
}
