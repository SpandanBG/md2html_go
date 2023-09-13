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
	italicMarker                = '*'
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

	stack := make([]int, 2)
	si := 0
	skip := false
	for i, char := range rawMD {
		if skip {
			skip = false
			continue
		}

		if char == common.EscChar {
			skip = true
			continue
		}

		if char == italicMarker {
			stack[si] = i
			si++
		}

		if si == 2 {
			si = 0
			txtRanges = append(txtRanges, common.TextRange{
				Range: []int{stack[0], stack[1] + 1},
				Type:  ItalicMarker,
			})
		}
	}

	return txtRanges
}
