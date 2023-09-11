package paragraph

import (
	"strings"

	"sudocoding.xyz/md2html_go/common"
	"sudocoding.xyz/md2html_go/richtext"
)

type Paragraph struct {
	Components []common.MDComponent
}

func ExtractParagraph(rawMD string) Paragraph {
	rt := richtext.ExtractRichText(rawMD)

	return Paragraph{
		Components: []common.MDComponent{
			&rt,
		},
	}
}

func (p *Paragraph) ToHTMLString() string {
	var s strings.Builder

	for _, comp := range p.Components {
		s.WriteString(comp.ToHTMLString())
	}

	return s.String()
}
