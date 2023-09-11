package common

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
