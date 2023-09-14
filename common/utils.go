package common

import "strings"

func EscapeRawToHTML(char rune) []rune {
	switch char {
	case '&':
		return []rune{'&', 'a', 'm', 'p', ';'}
	case '"':
		return []rune{'&', 'q', 'u', 'o', 't', ';'}
	case '\'':
		return []rune{'&', '#', '3', '9', ';'}
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
// (U+000A) or carriage rturn (U+000D), followed by a `line_ending`
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
