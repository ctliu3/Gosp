package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// ((⟨variable1⟩ ⟨init1⟩) ...)

type Binds struct {
  Bindings []Bind
}

func NewBinds(binds []Bind) *Binds {
  return &Binds{binds}
}

func (self *Binds) Type() string {
  return const_.BINDS
}

func (self *Binds) Eval(env *scope.Scope) value.Value {
  return nil
}

func (self *Binds) String() string {
  return "binds"
}
