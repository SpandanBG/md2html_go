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

func TestToHTMLString(t *testing.T) {
	para := Paragraph{}

	htmlStr := para.ToHTMLString()

	isEqual[string](t, htmlStr, "")
}

func TestParagraphMD(t *testing.T) {
	mdStr := "hello, world"
	para := ExtractParagraph(mdStr)

	htmlStr := para.ToHTMLString()

	isEqual[string](t, htmlStr, mdStr)
}
