package italic

import (
	"fmt"
	"strings"

	"sudocoding.xyz/md2html_go/common"
)

const (
	openingTag         = "<em>"
	closingTag         = "</em>"
	defaultPlaceholder = "%s%s%s"
)

type Italic struct {
	Value strings.Builder
}

func ExtractItalic(rawMD string) Italic {
	var s strings.Builder

	for i, char := range rawMD {
		if i == 0 || i == len(rawMD)-1 {
			continue
		}

		for _, escChar := range common.EscapeRawToHTML(char) {
			s.WriteRune(escChar)
		}
	}

	return Italic{Value: s}
}

func (i *Italic) ToHTMLString() string {
	return fmt.Sprintf(defaultPlaceholder, openingTag, i.Value.String(), closingTag)
}
