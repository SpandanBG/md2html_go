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
		name   string
		input  string
		output string
	}{
		{
			name:   "should change '' to '<em></em>'",
			input:  "",
			output: "<em></em>",
		},
		{
			name:   "should change '**' to '<em></em>'",
			input:  "**",
			output: "<em></em>",
		},
		{
			name:   "should change '*hello, world*' to '<em>hello, world</em>'",
			input:  "*hello, world*",
			output: "<em>hello, world</em>",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			actualOutput := ExtractItalic(test.input)

			isEqual[string](t, actualOutput.ToHTMLString(), test.output)
		})
	}
}
