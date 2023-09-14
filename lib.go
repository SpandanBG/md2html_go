package md2htmlgo

import (
	"sudocoding.xyz/md2html_go/common"
)

/* PIKE MATCHBOX */

func NewMarkdown(content string) common.MDRanges {
	return common.MDRanges{
		RawMD:        []rune(content),
		StartIndices: []int{},
		EndIndices:   []int{},
		MDTokens:     []rune{},
	}
}
