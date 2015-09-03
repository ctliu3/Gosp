package eval

import (
  "fmt"
  "github.com/ctliu3/gosp/parser"
  "github.com/ctliu3/gosp/scope"
)

func Eval(expr string) string {
  nodes := parser.ParseFromString(expr)
  fmt.Printf("#ast = %d\n", len(nodes))

  env := scope.NewScope(nil)
  for _, node := range nodes {
    fmt.Printf("%s\n", node.Eval(env))
  }
  return ""
}
