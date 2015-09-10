package eval

import (
  "fmt"
  "strings"

  "github.com/ctliu3/gosp/parser"
  "github.com/ctliu3/gosp/scope"
)

func Eval(expr string) string {
  nodes := parser.ParseFromString(expr)
  fmt.Printf("#AST = %d\n", len(nodes))

  env := scope.NewScope(nil)
  var result []string
  for _, node := range nodes {
    if val := node.Eval(env); val != nil {
      result = append(result, val.String())
    }
  }

  return strings.Join(result, " ")
}
