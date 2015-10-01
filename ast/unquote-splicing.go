package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

type UnQuoteSplicing struct {
  expr Node
}

func NewUnQuoteSplicing(expr Node) *UnQuoteSplicing {
  return &UnQuoteSplicing{expr}
}

func (self *UnQuoteSplicing) Type() string {
  return const_.UNQUOTE_SPLICING
}

func (self *UnQuoteSplicing) Eval(env *scope.Scope) value.Value {
  return self.expr.Eval(env)
}

func (self *UnQuoteSplicing) String() string {
  return self.ExtRep()
}

func (self *UnQuoteSplicing) ExtRep() string {
  return self.expr.ExtRep()
}
