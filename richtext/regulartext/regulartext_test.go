package regulartext

import (
	"testing"

	"sudocoding.xyz/md2html_go/common"
)

func isEqual[T comparable](t *testing.T, a, b T) {
	if a != b {
		t.Fatal("LHS: ", a, " != RHS: ", b)
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

func TestNewRegularText(t *testing.T) {
	for _, test := range []struct {
		name     string
		input    string
		output   string
		hasError bool
	}{
		{
			name:     "should return error for empty string",
			input:    "",
			output:   "",
			hasError: true,
		},
		{
			name:     "should return passed string as output",
			input:    "hello, world",
			output:   "hello, world",
			hasError: false,
		},
		{
			name:     "should escape chars passed",
			input:    "&\"'<>",
			output:   "&amp;&quot;&#39;&lt;&gt;",
			hasError: false,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			actualOuput, err := NewRegularText(test.input)

			if test.hasError {
				isNotNil[error](t, err)
			} else {
				isNil[error](t, err)
				isEqual[string](t, actualOuput.ToHTMLString(), test.output)
			}
		})
	}
}

func TestFillEmptyRanges(t *testing.T) {
	for _, test := range []struct {
		name      string
		input     []common.TextRange
		lastIndex int
		ouput     []common.TextRange
	}{
		{
			name:      "should fill entire ranges",
			input:     []common.TextRange{},
			lastIndex: 7,
			ouput: []common.TextRange{
				{Range: []int{0, 7}, Type: RegularTextMarker},
			},
		},
		{
			name: "should fill no ranges",
			input: []common.TextRange{
				{Range: []int{0, 7}, Type: "xyz"},
			},
			lastIndex: 7,
			ouput: []common.TextRange{
				{Range: []int{0, 7}, Type: "xyz"},
			},
		},
		{
			name: "should fill starting range",
			input: []common.TextRange{
				{Range: []int{4, 7}, Type: "xyz"},
			},
			lastIndex: 7,
			ouput: []common.TextRange{
				{Range: []int{0, 4}, Type: RegularTextMarker},
				{Range: []int{4, 7}, Type: "xyz"},
			},
		},
		{
			name: "should fill end range",
			input: []common.TextRange{
				{Range: []int{0, 3}, Type: "xyz"},
			},
			lastIndex: 7,
			ouput: []common.TextRange{
				{Range: []int{0, 3}, Type: "xyz"},
				{Range: []int{3, 7}, Type: RegularTextMarker},
			},
		},
		{
			name: "should fill start and end range",
			input: []common.TextRange{
				{Range: []int{2, 3}, Type: "xyz"},
				{Range: []int{3, 5}, Type: "xyz"},
			},
			lastIndex: 7,
			ouput: []common.TextRange{
				{Range: []int{0, 2}, Type: RegularTextMarker},
				{Range: []int{2, 3}, Type: "xyz"},
				{Range: []int{3, 5}, Type: "xyz"},
				{Range: []int{5, 7}, Type: RegularTextMarker},
			},
		},
		{
			name: "should fill start, middle and end range",
			input: []common.TextRange{
				{Range: []int{2, 3}, Type: "xyz"},
				{Range: []int{5, 6}, Type: "xyz"},
			},
			lastIndex: 7,
			ouput: []common.TextRange{
				{Range: []int{0, 2}, Type: RegularTextMarker},
				{Range: []int{2, 3}, Type: "xyz"},
				{Range: []int{3, 5}, Type: RegularTextMarker},
				{Range: []int{5, 6}, Type: "xyz"},
				{Range: []int{6, 7}, Type: RegularTextMarker},
			},
		},
		{
			name: "should fill start and middle range",
			input: []common.TextRange{
				{Range: []int{2, 3}, Type: "xyz"},
				{Range: []int{5, 7}, Type: "xyz"},
			},
			lastIndex: 7,
			ouput: []common.TextRange{
				{Range: []int{0, 2}, Type: RegularTextMarker},
				{Range: []int{2, 3}, Type: "xyz"},
				{Range: []int{3, 5}, Type: RegularTextMarker},
				{Range: []int{5, 7}, Type: "xyz"},
			},
		},
		{
			name: "should fill middle and end range",
			input: []common.TextRange{
				{Range: []int{0, 3}, Type: "xyz"},
				{Range: []int{5, 6}, Type: "xyz"},
			},
			lastIndex: 7,
			ouput: []common.TextRange{
				{Range: []int{0, 3}, Type: "xyz"},
				{Range: []int{3, 5}, Type: RegularTextMarker},
				{Range: []int{5, 6}, Type: "xyz"},
				{Range: []int{6, 7}, Type: RegularTextMarker},
			},
		},
		{
			name: "should fill middle range",
			input: []common.TextRange{
				{Range: []int{0, 3}, Type: "xyz"},
				{Range: []int{5, 7}, Type: "xyz"},
			},
			lastIndex: 7,
			ouput: []common.TextRange{
				{Range: []int{0, 3}, Type: "xyz"},
				{Range: []int{3, 5}, Type: RegularTextMarker},
				{Range: []int{5, 7}, Type: "xyz"},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			actualOuput := FillEmptyRanges(test.input, test.lastIndex)

			isEqual[int](t, len(actualOuput), len(test.ouput))
			for i, rng := range actualOuput {
				isEqual[int](t, rng.Range[0], test.ouput[i].Range[0])
				isEqual[int](t, rng.Range[1], test.ouput[i].Range[1])
				isEqual[string](t, rng.Type, test.ouput[i].Type)
			}
		})
	}
}
