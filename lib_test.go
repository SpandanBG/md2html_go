package md2htmlgo

import (
	"testing"

	"sudocoding.xyz/md2html_go/common"
)

func isEqual[T comparable](t *testing.T, a, b T) {
	if a != b {
		t.Fatal("LHS: \"", a, "\" != RHS: \"", b, "\"")
	}
}

func isNil[T comparable](t *testing.T, a T) {
	var nilVaue T
	if a != nilVaue {
		t.Fatal(a, " is not nil ")
	}
}

func isNotNil[T comparable](t *testing.T, a T) {
	var nilVaue T
	if a == nilVaue {
		t.Fatal(a, " is not nil ")
	}
}

func TestNewMarkdown(t *testing.T) {
	for _, test := range []struct {
		name   string
		input  string
		output common.MDRanges
	}{
		{
			name:  "should create default md range",
			input: "hello, world",
			output: common.MDRanges{
				RawMD:        []rune("hello, world"),
				StartIndices: []int{},
				EndIndices:   []int{},
				MDTokens:     []rune{},
			},
		},
		{
			name:  "should create secured md range",
			input: string([]rune{'a', 'b', 'c', common.Null}),
			output: common.MDRanges{
				RawMD:        []rune{'a', 'b', 'c', common.ReplacementCha},
				StartIndices: []int{},
				EndIndices:   []int{},
				MDTokens:     []rune{},
			},
		},
		{
			name:  "should escape backslashed characters",
			input: string([]rune{'a', '\\', '&', common.Null}),
			output: common.MDRanges{
				RawMD:        []rune{'a', '&', 'a', 'm', 'p', ';', common.ReplacementCha},
				StartIndices: []int{},
				EndIndices:   []int{},
				MDTokens:     []rune{},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			actualOutput := NewMarkdown(test.input)

			isEqual[int](t, len(test.output.RawMD), len(actualOutput.RawMD))
			isEqual[int](t, len(test.output.StartIndices), len(actualOutput.StartIndices))
			isEqual[int](t, len(test.output.EndIndices), len(actualOutput.EndIndices))
			isEqual[int](t, len(test.output.MDTokens), len(actualOutput.MDTokens))

			isEqual[string](t, string(test.output.RawMD), string(actualOutput.RawMD))
			isEqual[string](t, string(test.output.MDTokens), string(actualOutput.MDTokens))

			for i, eachIdx := range actualOutput.StartIndices {
				isEqual[int](t, test.output.StartIndices[i], eachIdx)
			}

			for i, eachIdx := range actualOutput.EndIndices {
				isEqual[int](t, test.output.EndIndices[i], eachIdx)
			}
		})
	}
}
