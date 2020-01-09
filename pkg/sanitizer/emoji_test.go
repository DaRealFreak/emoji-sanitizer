package sanitizer

import (
	"testing"

	"github.com/DaRealFreak/emoji-sanitizer/pkg/sanitizer/options"
	"github.com/stretchr/testify/assert"
)

func TestNewSanitizer(t *testing.T) {
	sanitizer, err := NewSanitizer(VersionLatest, options.FailSilently(true), options.LoadFromOnline(false))
	assert.New(t).NoError(err)
	assert.New(t).NotNil(sanitizer)

	sanitizer, err = NewSanitizer(VersionLatest, options.FailSilently(true), options.LoadFromOnline(true))
	assert.New(t).NoError(err)
	assert.New(t).NotNil(sanitizer)
}
