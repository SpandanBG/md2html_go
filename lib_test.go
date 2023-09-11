package md2htmlgo

import (
	"os"
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

func TestNewMarkdown(t *testing.T) {
	testfileName := "examples/hello_world.md"
	expectedOutput, _ := os.ReadFile("examples/hello_world.html")
	expecteContent := strings.TrimSpace(string(expectedOutput))

	result, err := NewMarkdownFromFile(testfileName)

	isNil[error](t, err)
	isNotNil[*Markdown](t, &result)
	isEqual[string](t, result.ToHTMLString(), expecteContent)
}
