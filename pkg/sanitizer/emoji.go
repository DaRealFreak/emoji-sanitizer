// Package sanitizer provides a unicode emoji sanitizer
package sanitizer

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/DaRealFreak/emoji-sanitizer/pkg/sanitizer/options"
)

const (
	// Version10 is the path segment for version 1.0
	Version10 = "1.0"
	// Version20 is the path segment for version 2.0
	Version20 = "2.0"
	// Version30 is the path segment for version 3.0
	Version30 = "3.0"
	// Version40 is the path segment for version 4.0
	Version40 = "4.0"
	// Version50 is the path segment for version 5.0
	Version50 = "5.0"
	// Version110 is the path segment for version 11.0
	Version110 = "11.0"
	// Version120 is the path segment for version 12.0
	Version120 = "12.0"
	// Version121 is the path segment for version 12.1
	Version121 = "12.1"
	// VersionLatest is the path segment for the latest version
	VersionLatest = "latest"

	versionLatestOffline = Version121

	// EmojiDataURLPath is the URL path to load the emoji-data from when choosing online mode
	EmojiDataURLPath = "https://unicode.org/Public/emoji/%s/emoji-data.txt"
)

// Sanitizer provides an option to sanitize unicode emoji runes based on the version and options
type Sanitizer struct {
	options       []options.Option
	regexpPattern *regexp.Regexp
}

// NewSanitizer initializes the sanitizer, loads the unicode data and applies the options
func NewSanitizer(sanitizerOptions ...options.Option) (*Sanitizer, error) {
	sanitizer := &Sanitizer{
		options: sanitizerOptions,
	}

	// if no version is defined we set the version option to the latest version
	if version := sanitizer.getOption(options.UnicodeVersion("")); version == nil {
		sanitizer.options = append(sanitizerOptions, options.UnicodeVersion(VersionLatest))
	}

	if err := sanitizer.loadUnicodeEmojiPattern(); err != nil {
		return nil, err
	}

	return sanitizer, nil
}

func (s *Sanitizer) getUnicodeVersion() string {
	return fmt.Sprintf("%v", s.getOption(options.UnicodeVersion(VersionLatest)).GetValue())
}

func (s *Sanitizer) getOption(option options.Option) options.Option {
	for _, setOptions := range s.options {
		if fmt.Sprintf("%T", option) == fmt.Sprintf("%T", setOptions) {
			return setOptions
		}
	}

	return nil
}

func (s *Sanitizer) isOptionSet(option options.Option) bool {
	for _, setOptions := range s.options {
		if fmt.Sprintf("%T", option) == fmt.Sprintf("%T", setOptions) {
			return option.GetValue() == setOptions.GetValue()
		}
	}

	return false
}

func (s *Sanitizer) loadUnicodeEmojiPattern() error {
	content, err := s.getEmojiDataContent()
	if err != nil {
		return err
	}

	var emojiUnicodeValues []string

	// match [4 bytes] and [4 bytes .. 4 bytes]
	unicodeEmojiLines := regexp.MustCompile(`(?m)^([0-9A-F]{4,5}(\.\.[0-9A-F]{4,5})?)\s+;`)

	for _, line := range strings.Split(string(content), "\n") {
		matches := unicodeEmojiLines.FindStringSubmatch(line)
		if len(matches) > 1 && matches[1] != "" {
			if !s.isEmojiCodeAllowed(matches[1]) {
				emojiUnicodeValues = append(emojiUnicodeValues, matches[1])
			}
		}
	}

	var emojiUnicodeRegexPattern string

	for _, emojiUnicode := range emojiUnicodeValues {
		if strings.Contains(emojiUnicode, "..") {
			emojiUnicodeRange := strings.Split(emojiUnicode, "..")
			emojiUnicodeRegexPattern += fmt.Sprintf(`\x{%s}-\x{%s}`, emojiUnicodeRange[0], emojiUnicodeRange[1])
		} else {
			emojiUnicodeRegexPattern += fmt.Sprintf(`\x{%s}`, emojiUnicode)
		}
	}

	emojiUnicodeRegexPattern = fmt.Sprintf(`[%s]`, emojiUnicodeRegexPattern)
	s.regexpPattern, err = regexp.Compile(emojiUnicodeRegexPattern)

	return err
}

func (s *Sanitizer) getEmojiDataContent() ([]byte, error) {
	if customPath := s.getOption(options.LoadFromCustomPath("")); customPath != nil {
		u, err := url.Parse(fmt.Sprintf("%v", customPath.GetValue()))
		if err != nil {
			return nil, err
		}

		var (
			r    io.ReadCloser
			resp *http.Response
		)

		// If it's a URL
		if u.Scheme == "http" || u.Scheme == "https" {
			resp, err = http.Get(fmt.Sprintf("%v", customPath.GetValue()))
			if err != nil {
				return nil, err
			}

			r = resp.Body
		} else {
			r, err = os.Open(fmt.Sprintf("%v", customPath.GetValue()))
			if err != nil {
				return nil, err
			}
		}

		return ioutil.ReadAll(r)
	}

	if s.isOptionSet(options.LoadFromOnline(true)) {
		emojiURL := fmt.Sprintf(EmojiDataURLPath, s.getUnicodeVersion())

		// #nosec G107
		res, err := http.Get(emojiURL)
		if err != nil {
			return nil, err
		}

		return ioutil.ReadAll(res.Body)
	}

	version := s.getUnicodeVersion()
	if version == VersionLatest {
		version = versionLatestOffline
	}

	return ioutil.ReadFile(fmt.Sprintf("emoji_data/%s/emoji-data.txt", version))
}

// isEmojiCodeAllowed checks the whitelist for allowed unicode emojis
func (s *Sanitizer) isEmojiCodeAllowed(unicodeCode string) bool {
	if codes := s.getOption(options.AllowEmojiCodes([]string{})); codes != nil {
		if w, ok := codes.GetValue().([]string); ok {
			for _, allowedUnicodeCode := range w {
				if allowedUnicodeCode == unicodeCode {
					return true
				}
			}
		}
	}

	return false
}

// ReplaceUnicodeEmojis replaces all unicode emojis from the passed subject with the passed replacement
func (s *Sanitizer) ReplaceUnicodeEmojis(subject string, repl string) string {
	return s.regexpPattern.ReplaceAllString(subject, repl)
}

// StripUnicodeEmojis strips all unicode emojis from the passed subject
func (s *Sanitizer) StripUnicodeEmojis(subject string) string {
	return s.ReplaceUnicodeEmojis(subject, "")
}
