package common

import (
	"fmt"
	"strings"
)

type TextRange struct {
	Range []int
	Type  MDMarkers
}

type RawText string

func (rt *RawText) ToHTMLString() string {
	return string(*rt)
}

type TaggedText struct {
	Components []MDComponent
	OpenTag    string
	CloseTag   string
}

func (tt *TaggedText) ToHTMLString() string {
	var s strings.Builder

	for _, comp := range tt.Components {
		s.WriteString(comp.ToHTMLString())
	}

	return fmt.Sprintf("%s%s%s", tt.OpenTag, s.String(), tt.CloseTag)
}
