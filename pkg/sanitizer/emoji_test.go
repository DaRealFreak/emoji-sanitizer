package sanitizer

import (
	"testing"

	"github.com/DaRealFreak/emoji-sanitizer/pkg/sanitizer/options"
	"github.com/stretchr/testify/assert"
)

func TestNewSanitizer(t *testing.T) {
	sanitizer, err := NewSanitizer(VersionLatest, options.LoadFromOnline(false))
	assert.New(t).NoError(err)
	assert.New(t).NotNil(sanitizer)

	sanitizer, err = NewSanitizer(VersionLatest, options.LoadFromOnline(true))
	assert.New(t).NoError(err)
	assert.New(t).NotNil(sanitizer)

	// general emoji codes which are normally allowed in most contexts
	// "#", "*", "[0-9]", "©", "®", "‼", "™"
	options.AllowEmojiCodes([]string{"0023", "002A", "0030..0039", "00A9", "00AE", "203C", "2122"})
}
