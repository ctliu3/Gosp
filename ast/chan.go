package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

type Chan struct {
  size int
}

func NewChan(size int) *Chan {
  return &Chan{size: size}
}

func (self *Chan) Type() string {
  return const_.CHAN
}

func (self *Chan) Eval(env *scope.Scope) value.Value {
  return value.NewChan(self.size)
}

func (self *Chan) String() string {
  return ""
}

func (self *Chan) ExtRep() string {
  return ""
}
