package regulartext

import (
	"strings"

	"sudocoding.xyz/md2html_go/common"
)

type RegularText struct {
	Value strings.Builder
}

func ExtractRegularText(rawMD string) RegularText {
	var s strings.Builder

	for _, char := range rawMD {
		for _, escChar := range common.EscapeRawToHTML(char) {
			s.WriteRune(escChar)
		}
	}

	return RegularText{Value: s}
}

func (rt *RegularText) ToHTMLString() string {
	return rt.Value.String()
}
