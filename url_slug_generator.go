package url_slug_generator

import (
	"strings"
	"unicode"
)

// Slugger converts arbitrary UTF-8 text into URL-safe ASCII slugs.
//
// A slug is defined here as a string of lowercase ASCII letters [a-z],
// digits [0-9], and single separator characters. Non-ASCII characters
// are dropped entirely rather than romanized; this keeps the dependency
// surface at zero (standard library only) and the output deterministic.
// Consecutive disallowed characters collapse into a single separator,
// and leading/trailing separators are removed.
type Slugger struct {
	Separator string
}

// New returns a Slugger using the given separator.
// The separator itself must be ASCII-safe for a meaningful URL slug;
// callers are responsible for choosing one (commonly "-" or "_").
// If an empty separator is given, "-" is used as a safe default so
// that the output is always a contiguous string rather than a run of
// merged words.
func New(separator string) *Slugger {
	if separator == "" {
		separator = "-"
	}
	return &Slugger{Separator: separator}
}

// Generate produces a slug from s.
func (sl *Slugger) Generate(s string) string {
	sep := sl.Separator
	var b strings.Builder
	prevSep := true

	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z':
			b.WriteRune(r)
			prevSep = false
		case r >= 'A' && r <= 'Z':
			b.WriteRune(unicode.ToLower(r))
			prevSep = false
		case r >= '0' && r <= '9':
			b.WriteRune(r)
			prevSep = false
		default:
			// Collapse consecutive disallowed characters into one separator.
			if !prevSep {
				b.WriteString(sep)
				prevSep = true
			}
		}
	}

	result := b.String()
	result = strings.TrimSuffix(result, sep)
	return result
}

// Generate is a package-level convenience using the default separator "-".
func Generate(s string) string {
	return New("").Generate(s)
}
