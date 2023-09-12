package richtext

import (
	"sudocoding.xyz/md2html_go/common"
	"sudocoding.xyz/md2html_go/richtext/italic"
	"sudocoding.xyz/md2html_go/richtext/regulartext"
)

func NewRichText(rawMD string) common.MDComponent {
	components := []common.MDComponent{}

	italicRanges := italic.GetItalicRanges(rawMD)
	allRanges := regulartext.FillEmptyRanges(italicRanges, len(rawMD))

	for _, rng := range allRanges {
		switch rng.Type {
		case italic.ItalicMarker:
			if italicText, err := italic.NewItalic(rawMD[rng.Range[0]:rng.Range[1]]); err == nil {
				components = append(components, italicText)
			}
		default:
			if regularText, err := regulartext.NewRegularText(rawMD[rng.Range[0]:rng.Range[1]]); err == nil {
				components = append(components, regularText)
			}
		}
	}

	return &common.TaggedText{
		Components: components,
		OpenTag:    "",
		CloseTag:   "",
	}
}
