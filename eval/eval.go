package eval

import (
  "fmt"
  "strings"

  "github.com/ctliu3/gosp/parser"
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/ast"
)

func Eval(expr string) string {
  nodes := parser.ParseFromString(expr)
  fmt.Printf("#AST = %d\n", len(nodes))

  env := scope.NewScope(nil)
  var result []string
  for _, node := range nodes {
    switch node.(type) {
    case *ast.Import:
      expr := node.Eval(env)
      EvalImport(expr.String(), env)
    default:
      if val := node.Eval(env); val != nil {
        result = append(result, val.String())
      }
    }
  }

  return strings.Join(result, " ")
}

func EvalImport(expr string, env *scope.Scope) {
  nodes := parser.ParseFromString(expr)

  for _, node := range nodes {
    node.Eval(env)
  }
}
