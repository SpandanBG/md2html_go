package md2htmlgo

import (
	"sudocoding.xyz/md2html_go/common"
	"sudocoding.xyz/md2html_go/paragraph"
)

func NewMarkdown(content string) common.MDComponent {
	para := paragraph.NewParagraph(content)

	return &common.TaggedText{
		Components: []common.MDComponent{
			para,
		},
		OpenTag:  "",
		CloseTag: "",
	}
}
