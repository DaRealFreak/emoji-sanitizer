package sanitizer

import (
	"testing"

	"github.com/DaRealFreak/emoji-sanitizer/pkg/sanitizer/options"
	"github.com/stretchr/testify/assert"
)

func TestNewSanitizer(t *testing.T) {
	// load the emoji data from offline
	sanitizer, err := NewSanitizer(options.LoadFromOnline(false))
	assert.New(t).NoError(err)
	assert.New(t).NotNil(sanitizer)
	assert.New(t).Equal(
		"Test string  ",
		sanitizer.StripUnicodeEmojis("Test string 😆😆😆 😛#123"),
	)

	// load the emoji data from online (https://unicode.org/)
	sanitizer, err = NewSanitizer(
		options.LoadFromOnline(true),
		options.UseFallbackToOffline(true),
	)
	assert.New(t).NoError(err)
	assert.New(t).NotNil(sanitizer)
	assert.New(t).Equal(
		"Test string  ",
		sanitizer.StripUnicodeEmojis("Test string 😆😆😆 😛#123"),
	)

	// load from existing custom online path (https://unicode.org/)
	sanitizer, err = NewSanitizer(options.LoadFromCustomPath("https://unicode.org/Public/emoji/latest/emoji-data.txt"))
	assert.New(t).NoError(err)
	assert.New(t).NotNil(sanitizer)
	assert.New(t).Equal(
		"Test string  ",
		sanitizer.StripUnicodeEmojis("Test string 😆😆😆 😛#123"),
	)

	// load from not existing custom online path
	sanitizer, err = NewSanitizer(options.LoadFromCustomPath("https://unicode.org/Public/emoji/nope/emoji-data.txt"))
	assert.New(t).Error(err)
	assert.New(t).Nil(sanitizer)

	// load from existing custom offline path
	sanitizer, err = NewSanitizer(options.LoadFromCustomPath("emoji_data/12.1/emoji-data.txt"))
	assert.New(t).NoError(err)
	assert.New(t).NotNil(sanitizer)
	assert.New(t).Equal(
		"Test string  ",
		sanitizer.StripUnicodeEmojis("Test string 😆😆😆 😛#123"),
	)

	// load from not existing custom offline path
	sanitizer, err = NewSanitizer(options.LoadFromCustomPath("emoji_data/not-existing-version/emoji-data.txt"))
	assert.New(t).Error(err)
	assert.New(t).Nil(sanitizer)

	sanitizer, err = NewSanitizer(
		options.UnicodeVersion(VersionLatest),
		// use offline data
		options.LoadFromOnline(false),
		// general emoji codes which are normally allowed in most contexts
		// "#", "*", "[0-9]", "©", "®", "‼", "™"
		options.AllowEmojiCodes([]string{"0023", "002A", "0030..0039", "00A9", "00AE", "203C", "2122"}),
	)
	assert.New(t).NoError(err)
	assert.New(t).NotNil(sanitizer)
	assert.New(t).Equal("Test string  #123", sanitizer.StripUnicodeEmojis("Test string 😆😆😆 😛#123"))
}
