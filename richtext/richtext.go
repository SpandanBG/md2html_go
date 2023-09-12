package richtext

import (
	"sudocoding.xyz/md2html_go/common"
	"sudocoding.xyz/md2html_go/richtext/regulartext"
)

func NewRichText(rawMD string) common.MDComponent {
	regulartext := regulartext.NewRegularText(rawMD)

	return &common.TaggedText{
		Components: []common.MDComponent{
			regulartext,
		},
		OpenTag:  "",
		CloseTag: "",
	}
}
