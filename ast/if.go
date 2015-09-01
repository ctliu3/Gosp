package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

type If struct {
  test Node
  conseq Node
  alt Node
}

func NewIf(test Node, conseq Node, alt Node) *If {
  return &If{test, conseq, alt}
}

func (self *If) Type() string {
  return "If"
}

func (self *If) Eval(env *scope.Scope) value.Value {
  res := self.test.Eval(env)
  if res.String() == const_.TRUE {
    return self.conseq.Eval(env)
  } else {
    return self.alt.Eval(env)
  }
}

func (self *If) String() string {
  return "if"
}
