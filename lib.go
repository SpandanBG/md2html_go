package md2htmlgo

import (
	"sudocoding.xyz/md2html_go/common"
)

/* PIKE MATCHBOX */

func NewMarkdown(content string) common.MDRanges {
	cleanedContent := common.SecureNullChar(content)

	// Call escape backslashed on content only after all MD meaning have been
	// translated.
	// This is to avoid incorrect escaping in strings like: "\\`test`"
	// This should produce "\<code>test</code>" and not "`test`"
	escapedContent := common.EscBackslashed(cleanedContent)

	return common.MDRanges{
		RawMD:        []rune(escapedContent),
		StartIndices: []int{},
		EndIndices:   []int{},
		MDTokens:     []rune{},
	}
}
