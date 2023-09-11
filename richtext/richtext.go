package richtext

import (
	"strings"

	"sudocoding.xyz/md2html_go/common"
	"sudocoding.xyz/md2html_go/richtext/regulartext"
)

type RichText struct {
	Components []common.MDComponent
}

func ExtractRichText(rawMD string) RichText {
	regulartext := regulartext.ExtractRegularText(rawMD)

	return RichText{
		Components: []common.MDComponent{
			&regulartext,
		},
	}
}

func (rt *RichText) ToHTMLString() string {
	var s strings.Builder
	for _, comp := range rt.Components {
		s.WriteString(comp.ToHTMLString())
	}
	return s.String()
}
