package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// (⟨variable1⟩ ⟨init1⟩)

type Bind struct {
  var_ string
  init Node
}

func NewBind(var_ string, init Node) *Bind {
  return &Bind{var_, init}
}

func (self *Bind) Type() string {
  return const_.BIND
}

func (self *Bind) Eval(env *scope.Scope) value.Value {
  return nil
}

func (self *Bind) String() string {
  return "bind"
}

func (self *Bind) ExtRep() string {
  return ""
}
