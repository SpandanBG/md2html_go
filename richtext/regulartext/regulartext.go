package regulartext

import (
	"errors"
	"strings"

	"sudocoding.xyz/md2html_go/common"
)

const RegularTextMarker = "aA"

const noTextErrMsg = "No text present to create regular text MD component"

func NewRegularText(rawMD string) (common.MDComponent, error) {
	if rawMD == "" {
		return nil, errors.New(noTextErrMsg)
	}

	var s strings.Builder
	for _, char := range rawMD {
		for _, escChar := range common.EscapeRawToHTML(char) {
			s.WriteRune(escChar)
		}
	}
	comp := common.RawText(s.String())
	return &comp, nil
}

func FillEmptyRanges(ranges []common.TextRange, lastIndex int) []common.TextRange {
	if len(ranges) == 0 {
		return []common.TextRange{{
			Range: []int{0, lastIndex},
			Type:  RegularTextMarker,
		}}
	}

	allRanges := []common.TextRange{}

	i := 0
	for j := 0; j < len(ranges); j++ {
		if i == ranges[j].Range[0] {
			allRanges = append(allRanges, ranges[j])
			i = ranges[j].Range[1]
			continue
		}

		allRanges = append(allRanges, common.TextRange{
			Range: []int{i, ranges[j].Range[0]},
			Type:  RegularTextMarker,
		}, ranges[j])
		i = ranges[j].Range[1]
	}

	if i < lastIndex {
		allRanges = append(allRanges, common.TextRange{
			Range: []int{i, lastIndex},
			Type:  RegularTextMarker,
		})
	}

	return allRanges
}
