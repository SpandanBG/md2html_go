package md2htmlgo

import (
	"sudocoding.xyz/md2html_go/common"
)

/* PIKE MATCHBOX */

func NewMarkdown(content string) common.MDRanges {
	cleanedContent := common.SecureNullChar(content)

	return common.MDRanges{
		RawMD:        []rune(cleanedContent),
		StartIndices: []int{},
		EndIndices:   []int{},
		MDTokens:     []rune{},
	}
}
