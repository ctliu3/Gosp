package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

type UnQuote struct {
  expr Node
}

func NewUnQuote(expr Node) *UnQuote {
  return &UnQuote{expr}
}

func (self *UnQuote) Type() string {
  return const_.UNQUOTE
}

func (self *UnQuote) Eval(env *scope.Scope) value.Value {
  return self.expr.Eval(env)
}

func (self *UnQuote) String() string {
  return self.ExtRep()
}

func (self *UnQuote) ExtRep() string {
  return self.String()
}
