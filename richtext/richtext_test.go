package richtext

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

func TestRichText(t *testing.T) {
	richText := RichText{}

	htmlStr := richText.ToHTMLString()

	isEqual[string](t, htmlStr, "")
}

func TextRichTextWithJustRegularText(t *testing.T) {
	mdStr := "hello, world"
	rt := ExtractRichText(mdStr)

	htmlStr := rt.ToHTMLString()

	isEqual[string](t, htmlStr, mdStr)
}
