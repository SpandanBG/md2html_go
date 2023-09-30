package common

import "strings"

// Escapes the escapable character to their respective escaped HTML equivalent
//
// This includes the characters [&, ", <, >]
func EscapeRawToHTML(char rune) []rune {
	switch char {
	case '&':
		return []rune{'&', 'a', 'm', 'p', ';'}
	case '"':
		return []rune{'&', 'q', 'u', 'o', 't', ';'}
	case '<':
		return []rune{'&', 'l', 't', ';'}
	case '>':
		return []rune{'&', 'g', 't', ';'}
	default:
		return []rune{char}
	}
}

// Split the string into multiple lines
//
// A `line` is a sequence of zero or more `characters` other than line feed
// (U+000A) or carriage return (U+000D), followed by a `line_ending`
//
// A `line_ending` is a line feed (U+000A), a carriage return (U+000A) followed
// with or without a line feed.
func SplitByLines(wholeContent string) []string {
	var eachLine strings.Builder
	lines := []string{}

	prevCharCarriageReturn := false
	for _, char := range wholeContent {
		if prevCharCarriageReturn {
			prevCharCarriageReturn = false
			if char == LineFeed {
				continue
			}
		}

		if char == LineFeed || char == CarriageReturn {
			lines = append(lines, eachLine.String())
			eachLine.Reset()
			prevCharCarriageReturn = char == CarriageReturn
			continue
		}

		eachLine.WriteRune(char)
	}

	lastLine := eachLine.String()
	if lastLine != "" {
		lines = append(lines, lastLine)
	}

	return lines
}

// Returns true if a line is a blank line by markdown specs
//
// A line containing no characters, or a line containing only spaces (U+0020)
// or tabs (U+0009), is a blank line
func IsBlankLine(line string) bool {
	if line == "" {
		return true
	}

	for _, char := range line {
		if char != Space && char != Tab {
			return false
		}
	}

	return true
}

// Returns true if the character is a whitespace character
//
// A Unicode whitespace character is any code point in the Unicode Zs category
// or a tab (U+0009), line feed (U+000A), form feed (U+000C), or carriage return
// (U+000D)
func IsCharWhiteSpace(char rune) bool {
	switch char {
	case Space, Tab, LineFeed, CarriageReturn:
		return true
	}

	if char >= ENQuad && char <= HairSpace {
		return true
	}

	switch char {
	case FormFeed, NoBreakSpace, OGHamSpaceMark, NarrowNoBreakSpace, MediumMathematicalSpace, IdeographicSpace:
		return true
	}

	return false
}

// Changes Null (U+0000) to Replacement Character (U+FFFD)
//
// Following the common markdown security guidelines
func SecureNullChar(content string) string {
	var s strings.Builder

	for _, char := range content {
		if char == Null {
			s.WriteRune(ReplacementCha)
		} else {
			s.WriteRune(char)
		}
	}

	return s.String()
}

// Returns `true` if the character is an ASCII Punctuation Character
//
// The ASCII ranges are from U+0021–2F, U+003A–0040, U+005B–0060, U+007B–007E
func IsASCIIPunctuationCharacter(char rune) bool {
	if char >= '\u0021' && char <= '\u002F' {
		return true
	}

	if char >= '\u003A' && char <= '\u0040' {
		return true
	}

	if char >= '\u005B' && char <= '\u0060' {
		return true
	}

	if char >= '\u007B' && char <= '\u007E' {
		return true
	}

	return false
}

// Escapes ASCII punctuations if backslash escaped
//
// The backslash esc character '\', escapes all ASCII punctuations
// Before other characters, the esc character is treated as literal backslashes
// Escaped characters are treated as regular characters and do not have their
// usual Markdown meanings
func EscBackslashed(line string) string {
	var s strings.Builder

	prevCharEscChar := false
	for _, char := range line {
		if char == EscChar && !prevCharEscChar {
			prevCharEscChar = true
			continue
		}

		if prevCharEscChar {
			if IsASCIIPunctuationCharacter(char) {
				s.WriteString(string(EscapeRawToHTML(char)))
			} else {
				s.WriteRune(EscChar)
				s.WriteRune(char)
			}
			prevCharEscChar = false
		} else {
			s.WriteRune(char)
		}
	}

	return s.String()
}
