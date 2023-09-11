package paragraph

import (
	"sudocoding.xyz/md2html_go/common"
)

type Paragraph struct {
	Components []common.MDComponent
}

func (p *Paragraph) ToHTMLString() string {
	return ""
}
