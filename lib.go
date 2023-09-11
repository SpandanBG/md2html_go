package md2htmlgo

import (
	"os"

	"sudocoding.xyz/md2html_go/common"
	"sudocoding.xyz/md2html_go/paragraph"
)

type Markdown struct {
	Components []common.MDComponent
}

func NewMarkdownFromFile(filename string) (Markdown, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return Markdown{}, err
	}

	content := string(file)

	para := paragraph.ExtractParagraph(content)

	return Markdown{
		Components: []common.MDComponent{
			&para,
		},
	}, nil
}
