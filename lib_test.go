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
	result := NewMarkdown("hello, world")

	isNotNil[common.MDComponent](t, result)
	isEqual[string](t, result.ToHTMLString(), "<p>hello, world</p>")
}
