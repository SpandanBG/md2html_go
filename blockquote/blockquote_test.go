package blockquote

import (
	"strings"
	"testing"
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

func TestGetBlockQuoteContent(t *testing.T) {
	for _, test := range []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "should return emtpy string for empty >",
			input:  ">",
			output: "",
		},
		{
			name:   "should return a for >a",
			input:  ">a",
			output: "a",
		},
		{
			name:   "should return a for > a",
			input:  "> a",
			output: "a",
		},
		{
			name:   "should return '  a' for >\ta",
			input:  ">\ta",
			output: "  a",
		},
		{
			name:   "should return '  \ta' for >\t\ta",
			input:  ">\t\ta",
			output: "  \ta",
		},
		{
			name:   "should return '' for >\t",
			input:  ">\t",
			output: "",
		},
		{
			name:   "should return '' for > followed by just spaces and tabs",
			input:  ">\t   \t   \t\t    \t",
			output: "",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			actualOutput := GetBlockQuoteContent(test.input)

			isEqual[string](t, test.output, actualOutput)
		})
	}
}

func BenchmarkGetBlockQuoteContent_AllWhiteSpaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetBlockQuoteContent(">" + strings.Repeat("  ", i))
	}
}
