package blockquote

import (
	"strings"

	"sudocoding.xyz/md2html_go/common"
)

// Returns the content of blockquote line. (Assumption is the string starts with
// `>` character).
//
// The `>` that beigns a block quote may be followed optionally by a space which
// isn't considered as part of the content.
// So, in case `>` is followed by a tab, it treated as 3 spaces
// That means 1st space will be discarded and the content will start with
// 2 spaces
func GetBlockQuoteContent(blockquoteLine string) string {
	// If just `>` is the entire line
	if len(blockquoteLine) == 1 {
		return ""
	}

	contentChars := []rune(blockquoteLine[1:])
	var s strings.Builder
	isNonEmptyCharWritten := false

	// Push 3 spaces to content string if first char is a Tab
	if contentChars[0] == common.Tab {
		// Early exit with empty if nothing else is present
		if len(contentChars) == 1 {
			return ""
		}

		s.WriteString(string([]rune{common.Space, common.Space}))

	} else if contentChars[0] != common.Space {
		s.WriteRune(contentChars[0])
		isNonEmptyCharWritten = true
	}

	for _, char := range contentChars[1:] {
		isNonEmptyCharWritten = char != common.Tab && char != common.Space
		s.WriteRune(char)
	}

	if !isNonEmptyCharWritten {
		return ""
	}

	return s.String()
}
