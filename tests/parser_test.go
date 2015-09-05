package tests

import (
  "fmt"
  "testing"
  "github.com/ctliu3/gosp/parser"
)

//func TestParser(t *testing.T) {
  //expr := `(if (> a 4) (display yes) (display no)))`
  //parser.ParseFromString(expr)
//}

func TestDefine(t *testing.T) {
  expr := `(define b 1)`
  nodes := parser.ParseFromString(expr)
  for _, node := range nodes {
    fmt.Println(node.Type())
  }
}
