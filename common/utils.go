package common

func EscAbleMarkes(char rune) bool {
	switch char {
	case rune(ItalicMarker):
		return true
	case EscChar:
		return true
	default:
		return false
	}
}

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
