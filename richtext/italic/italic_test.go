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
			name:     "should change '**' to '<em></em>'",
			input:    "**",
			output:   "<em></em>",
			hasError: false,
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
