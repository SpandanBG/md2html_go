package common

import "testing"

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

func TestEscapeRawToHTML(t *testing.T) {
	for _, test := range []struct {
		name   string
		input  rune
		output string
	}{
		{
			name:   "should turn & to &amp;",
			input:  '&',
			output: "&amp;",
		},
		{
			name:   "should turn \" to &quot;",
			input:  '"',
			output: "&quot;",
		},
		{
			name:   "should turn ' to &#39;",
			input:  '\'',
			output: "&#39;",
		},
		{
			name:   "should turn < to &lt;",
			input:  '<',
			output: "&lt;",
		},
		{
			name:   "should turn > to &gt;",
			input:  '>',
			output: "&gt;",
		},
		{
			name:   "should any other to itself",
			input:  'a',
			output: "a",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			actualOuput := EscapeRawToHTML(test.input)

			isEqual[string](t, test.output, string(actualOuput))
		})
	}
}
