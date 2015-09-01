package repl

import (
  "fmt"
  "github.com/ctliu3/gosp/parser"
)

func REPL(expr string) string {
  nodes := parser.ParseFromString(expr)
  fmt.Println(len(nodes))
  //for _, node := range nodes {
    //fmt.Printf("%s\n", node.String())
  //}
  return ""
}
