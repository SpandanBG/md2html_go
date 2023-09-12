package italic

import "testing"

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

func TestItalic(t *testing.T) {
	for _, test := range []struct {
		name     string
		input    string
		output   string
		hasError bool
	}{
		{
			name:     "should throw error for empty input",
			input:    "",
			output:   "",
			hasError: true,
		},
		{
			name:     "should throw error for single italics marker",
			input:    "*",
			output:   "",
			hasError: true,
		},
		{
			name:     "should throw error for empty italics content",
			input:    "**",
			output:   "",
			hasError: true,
		},
		{
			name:     "should throw error for more than 2 italics markers",
			input:    "*hello, *world*",
			output:   "",
			hasError: true,
		},
		{
			name:     "should change '*hello & world*' to '<em>hello &amp; world</em>'",
			input:    "*hello, world*",
			output:   "<em>hello, world</em>",
			hasError: false,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			actualOutput, err := NewItalic(test.input)

			if test.hasError {
				isNotNil[error](t, err)
			} else {
				isNil[error](t, err)
				isEqual[string](t, actualOutput.ToHTMLString(), test.output)
			}
		})
	}
}

func TestGetItalicRange(t *testing.T) {
	for _, test := range []struct {
		name          string
		input         string
		expectedRange [][]int
	}{
		{
			name:          "should return no range for emtpy string",
			input:         "",
			expectedRange: [][]int{},
		},
		{
			name:          "should return no range for no markers",
			input:         "hello, world",
			expectedRange: [][]int{},
		},
		{
			name:          "should return 1 range for 1 italic group",
			input:         "*a*",
			expectedRange: [][]int{{0, 3}},
		},
		{
			name:          "should return 2 range for 2 italic group",
			input:         "a *a* b *c*",
			expectedRange: [][]int{{2, 5}, {8, 11}},
		},
		{
			name:          "should return 2 range for 2 italic group and an extra *",
			input:         "a *a* *b *c*",
			expectedRange: [][]int{{2, 5}, {6, 10}},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			txtRanges := GetItalicRanges(test.input)

			isEqual[int](t, len(txtRanges), len(test.expectedRange))

			for i, rng := range txtRanges {
				isEqual[int](t, rng.Range[0], test.expectedRange[i][0])
				isEqual[int](t, rng.Range[1], test.expectedRange[i][1])
			}
		})
	}
}
