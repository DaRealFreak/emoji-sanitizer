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

type unicodeVersion string

// GetValue implements the option interface method
func (o unicodeVersion) GetValue() interface{} {
	return string(o)
}

// UnicodeVersion is the option to set a specific unicode version to load
// this option won't have any effect if you load the emoji data from a custom path
func UnicodeVersion(version string) Option {
	return unicodeVersion(version)
}

type offlineFallback bool

// GetValue implements the option interface method
func (o offlineFallback) GetValue() interface{} {
	return bool(o)
}

// UseFallbackToOffline is the option to fall back to loading the emoji data offline should an error occur
// this option won't have any effect if you don't use the LoadFromOnline option
func UseFallbackToOffline(useOfflineFallback bool) Option {
	return offlineFallback(useOfflineFallback)
}
