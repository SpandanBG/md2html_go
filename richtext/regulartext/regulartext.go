package regulartext

import (
	"strings"

	"sudocoding.xyz/md2html_go/common"
)

func NewRegularText(rawMD string) common.MDComponent {
	var s strings.Builder
	for _, char := range rawMD {
		for _, escChar := range common.EscapeRawToHTML(char) {
			s.WriteRune(escChar)
		}
	}
	comp := common.RawText(s.String())
	return &comp
}
