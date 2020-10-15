# Go Emoji Sanitizer
![tests](https://github.com/DaRealFreak/emoji-sanitizer/workflows/tests/badge.svg?branch=master) [![Coverage Status](https://coveralls.io/repos/github/DaRealFreak/emoji-sanitizer/badge.svg?branch=master)](https://coveralls.io/github/DaRealFreak/emoji-sanitizer?branch=master) ![GitHub](https://img.shields.io/github/license/DaRealFreak/emoji-sanitizer) [![Go Report Card](https://goreportcard.com/badge/github.com/DaRealFreak/emoji-sanitizer)](https://goreportcard.com/report/github.com/DaRealFreak/emoji-sanitizer)

go library to detect unicode emoji runes and to sanitize them.

## Usage
The usage is fairly simple. You generate a sanitizer and can sanitize passed strings.  
A minimalistic usage:
```go
sanitizer, err := NewSanitizer()
// this will return "Test string [e][e][e] [e]"
sanitizer.ReplaceUnicodeEmojis("Test string 😆😆😆 😛", "[e]")
// this will return: "Test string  "
sanitizer.StripUnicodeEmojis("Test string 😆😆😆 😛")
```

You can also set multiple options to further configure the sanitization.

## Options
### Set custom unicode version
You can simply set the unicode version with the UnicodeVersion option:
```go
sanitizer, err := NewSanitizer(
    options.UnicodeVersion(VersionLatest),
)
```

currently offline available are following versions:  
1.0, 2.0, 3.0, 4.0, 5.0, 11.0, 12.0, 12.1, 13.0, 13.1 
All of those are taken from https://unicode.org/Public/emoji/

If you are using the online option you can also instantly use other versions when they get released.

By default the version is set to the currently newest version: 13.1

### Emoji Data from custom path
You can load the emoji data from whatever path you want with the option:
```go
sanitizer, err := NewSanitizer(
    options.LoadFromCustomPath("https://unicode.org/Public/13.0.0/ucd/emoji/emoji-data.txt"),
)
```

If no `http` or `https` is in the scheme of the parsed URL it'll try to load a local file.

### Retrieve Emoji Data from offline/online
In case you always want to run the latest emoji data or possibly any updates you can use the option:
```go
// offline, this is already the default option
sanitizer, err := NewSanitizer(
    options.LoadFromOnline(false),
)

// online, you'll need this option if you want to load the emoji data from online
sanitizer, err := NewSanitizer(
    options.LoadFromOnline(true),
)
```

To reduce the error proneness you can further set the option to fallback to the offline approach should an error occur while loading the emoji data.
```go
sanitizer, err := NewSanitizer(
    options.LoadFromOnline(true),
    options.FallbackToOffline(true),
)
```

### Allow specific unicode emoji runes
In case you want to allow specific unicode emoji runes you can allow specific runes/ranges.  
Best shown in the example of basic runes which are categorized as emoji:  

```go
sanitizer, err := NewSanitizer(
    // general emoji codes which are normally allowed in most contexts
    // "#", "*", "[0-9]", "©", "®", "‼", "™"
    options.AllowEmojiCodes([]string{"0023", "002A", "0030..0039", "00A9", "00AE", "203C", "2122"}),
)
```

## Development
Want to contribute? Great!  
I'm always glad hearing about bugs or pull requests.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
