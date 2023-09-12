package paragraph

import (
	"sudocoding.xyz/md2html_go/common"
	"sudocoding.xyz/md2html_go/richtext"
)

const (
	openingTag = "<p>"
	closingTag = "</p>"
)

func NewParagraph(rawMD string) common.MDComponent {
	rt := richtext.NewRichText(rawMD)

	return &common.TaggedText{
		Components: []common.MDComponent{
			rt,
		},
		OpenTag:  openingTag,
		CloseTag: closingTag,
	}
}
