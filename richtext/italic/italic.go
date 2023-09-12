package italic

import (
	"errors"
	"regexp"

	"sudocoding.xyz/md2html_go/common"
	"sudocoding.xyz/md2html_go/richtext/regulartext"
)

const (
	ItalicMarker = "*"
)

const (
	openingTag                  = "<em>"
	closingTag                  = "</em>"
	validMDItalicREGXP          = "^\\*([^\\*]*)\\*$"
	searchItalicsREGXP          = "(?U)\\*.*\\*"
	invalidItalicMDStringErrMsg = "Invalid italic markdown supplied"
	emptyItalicsSuppliedErrMsg  = "Empty italic content supplied"
	noClosingMarkerErrMsg       = "No closing italic marker found"
)

func NewItalic(rawMD string) (common.MDComponent, error) {
	italicRexgp := regexp.MustCompile(validMDItalicREGXP)

	if !italicRexgp.Match([]byte(rawMD)) {
		return nil, errors.New(invalidItalicMDStringErrMsg)
	}

	secondLastIdx := len(rawMD) - 1
	regularText, err := regulartext.NewRegularText(rawMD[1:secondLastIdx])
	if err != nil {
		return nil, errors.New(emptyItalicsSuppliedErrMsg)
	}

	return &common.TaggedText{
		Components: []common.MDComponent{regularText},
		OpenTag:    openingTag,
		CloseTag:   closingTag,
	}, nil
}

func GetItalicRanges(rawMD string) []common.TextRange {
	txtRanges := []common.TextRange{}
	if rawMD == "" {
		return txtRanges
	}

	re := regexp.MustCompile(searchItalicsREGXP)
	slices := re.FindAllStringIndex(rawMD, -1)
	for _, rng := range slices {
		txtRanges = append(txtRanges, common.TextRange{
			Range: rng,
			Type:  ItalicMarker,
		})
	}
	return txtRanges
}
