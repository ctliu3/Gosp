package ast

import (
  "fmt"
  "time"
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

type Time struct {
  expr Node
}

func NewTime(expr Node) *Time {
  return &Time{expr}
}

func (self *Time) Type() string {
  return const_.TIME
}

func (self *Time) Eval(env *scope.Scope) value.Value {
  start := time.Now()
  ret := self.expr.Eval(env)
  elapsed := time.Since(start)
  fmt.Printf("Elapsed time: %.3f msecs\n", 1000.0 * elapsed.Seconds())
  return ret
}

func (self *Time) String() string {
  return "time"
}

func (self *Time) ExtRep() string {
  return ""
}
