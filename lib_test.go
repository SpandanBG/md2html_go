package md2htmlgo

import "testing"

func TestHello(t *testing.T)  {
  response := Hello()

  if response != "Hello, world" {
    t.Errorf("Hello, world error. Maybe become a better programmer")
  }
}
