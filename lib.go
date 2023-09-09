package md2htmlgo

import (
  "fmt"
)

func hello() string {
  return "Hello, world"
}

func ToParagraph(md_string string) string {
  return fmt.Sprintf("<p>%s</p>", md_string)
}
