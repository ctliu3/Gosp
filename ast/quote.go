package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

type Quote struct {
  value string
}

func NewQuote(name string) *Quote {
  return &Quote{value: name}
}

func (self *Quote) Type() string {
  return const_.QUOTE
}

func (self *Quote) Eval(env *scope.Scope) value.Value {
  return value.NewQuote(self.value)
}

func (self *Quote) String() string {
  return self.value
}
