package richtext

import (
	"sudocoding.xyz/md2html_go/common"
)

type RichText struct {
	Component []common.MDComponent
}

func (rt *RichText) ToHTMLString() string {
	return ""
}
