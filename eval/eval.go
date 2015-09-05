package eval

import (
  "fmt"
  "strings"

  "github.com/ctliu3/gosp/parser"
  "github.com/ctliu3/gosp/scope"
)

func Eval(expr string) string {
  nodes := parser.ParseFromString(expr)
  fmt.Printf("#ast = %d\n", len(nodes))

  env := scope.NewScope(nil)
  var res []string
  for _, node := range nodes {
    fmt.Println("#Eval(): " + node.Type())
    if val := node.Eval(env); val != nil {
      res = append(res, val.String())
    }
  }

  return strings.Join(res, " ")
}
