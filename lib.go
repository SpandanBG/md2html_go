package md2htmlgo

import (
	"sudocoding.xyz/md2html_go/common"
)

type Markdown struct {
	Components []common.MDComponent
}

func NewMarkdown(filename string) Markdown {
	return Markdown{}
}
