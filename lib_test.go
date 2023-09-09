package md2htmlgo

import "testing"

func isEqual[T comparable](t *testing.T, a, b T) {
  if a != b {
    t.Fatal("LHS: ", a, " != RHS: ", b)
  }
}

func TestHello(t *testing.T) {
  response := hello()

  isEqual[string](t, response, "Hello, world")
}

func TestParagraph(t *testing.T) {
  result := ToParagraph("hello, world")

  isEqual[string](t, result, "<p>hello, world</p>")
}
