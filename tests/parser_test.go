package tests

import (
  "fmt"
  "testing"
  "github.com/ctliu3/gosp/parser"
)

func parse(exprs []string) {
  for _, expr := range exprs {
    nodes := parser.ParseFromString(expr)
    for _, node := range nodes {
      fmt.Println(node.Type())
    }
  }
}

func TestParseDefine(t *testing.T) {
  exprs := []string {
    `(define b 1)`,
    `(define f (lambda (x y) (+ x y)))`,
  }
  parse(exprs)
}

func TestParseLambda(t *testing.T) {
  exprs := []string {
    `((lambda (x) (+ x x)) 4)`,
  }
  parse(exprs)
}

func TestParseLet(t *testing.T) {
  exprs := []string {
    `(let ((x 2) (y 3)) (+ x y))`,
  }
  parse(exprs)
}
