package tests

import (
  "testing"
  "github.com/ctliu3/gosp/parser"
)

func TestParser(t *testing.T) {
  expr := `(if (> a 4) (display yes) (display no)))`
  parser.ParseFromString(expr)
}
