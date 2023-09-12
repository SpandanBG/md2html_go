package paragraph

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

func TestParagraphMD(t *testing.T) {
	for _, test := range []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "should return <p>hello, world</p>",
			input:  "hello, world",
			output: "<p>hello, world</p>",
		},
		{
			name:   "should return <p>hello, <em>world</em></p>",
			input:  "hello, *world*",
			output: "<p>hello, <em>world</em></p>",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			para := NewParagraph(test.input)

			actualOutput := para.ToHTMLString()

			isEqual[string](t, actualOutput, test.output)
		})
	}
}
