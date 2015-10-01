package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

type Symbol struct {
  name string
}

func NewSymbol(name string) *Symbol {
  return &Symbol{name}
}

func (self *Symbol) Type() string {
  return const_.SYMBOL
}

func (self *Symbol) Eval(env *scope.Scope) value.Value {
  return value.NewSymbol(self.name)
}

func (self *Symbol) String() string {
  return self.name
}

func (self *Symbol) ExtRep() string {
  return self.name
}
