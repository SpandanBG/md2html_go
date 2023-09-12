package italic

import (
	"errors"
	"regexp"

	"sudocoding.xyz/md2html_go/common"
	"sudocoding.xyz/md2html_go/richtext/regulartext"
)

const (
	openingTag                  = "<em>"
	closingTag                  = "</em>"
	validMDItalicREXGP          = "^\\*([^\\*]*)\\*$"
	invalidItalicMDStringErrMsg = "Invalid italic markdown supplied"
)

func NewItalic(rawMD string) (common.MDComponent, error) {
	italicRexgp := regexp.MustCompile(validMDItalicREXGP)

	if !italicRexgp.Match([]byte(rawMD)) {
		return nil, errors.New(invalidItalicMDStringErrMsg)
	}

	secondLastIdx := len(rawMD) - 1
	regularText := regulartext.NewRegularText(rawMD[1:secondLastIdx])

	return &common.TaggedText{
		Components: []common.MDComponent{regularText},
		OpenTag:    openingTag,
		CloseTag:   closingTag,
	}, nil
}
