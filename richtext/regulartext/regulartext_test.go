package regulartext

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

func TestToHMTLString(t *testing.T) {
	mdStr := "hello, world"
	regulartext := ExtractRegularText(mdStr)

	htmlStr := regulartext.ToHTMLString()

	isEqual[string](t, htmlStr, mdStr)
}
