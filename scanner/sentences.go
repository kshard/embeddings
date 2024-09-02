package scanner

import (
	"bufio"
	"io"
	"unicode/utf8"
)

// Creates instance of [bufio.Scanner] configured for naïve sentence scanning.
func NewSentences(r io.Reader) *bufio.Scanner {
	s := bufio.NewScanner(r)
	s.Split(ScanSentence)
	return s
}

// ScanSentence is a split function for a [bufio.Scanner] that returns each
// sentence. It will never return an empty string.
// The definition of space is set by `[.!?]\s+|\z`
//
// Note: this is naïve algorithm primary used for testing purpose or simple text
func ScanSentence(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !isSpace(r) {
			break
		}
	}

	// Scan until end of sentence [.!?]\s+|\z
	var r rune
	for width, i := 0, start; i < len(data)-1; i += width {
		r, width = utf8.DecodeRune(data[i+1:])
		switch data[i] {
		case '.', '!', '?':
			if isSpace(r) {
				return i + 1, data[start : i+1], nil
			}
		}
	}

	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data[start:], nil
	}

	// Request more data.
	return 0, nil, nil
}

// isSpace reports whether the character is a Unicode white space character.
// We avoid dependency on the unicode package, but check validity of the implementation
// in the tests.
func isSpace(r rune) bool {
	if r <= '\u00FF' {
		// Obvious ASCII ones: \t through \r plus space. Plus two Latin-1 oddballs.
		switch r {
		case ' ', '\t', '\n', '\v', '\f', '\r':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}
	// High-valued ones.
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}
