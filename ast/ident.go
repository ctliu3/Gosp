package ast

import (
  //"fmt"
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
)

type Ident struct {
  name string
}

func NewIdent(name string) *Ident {
  return &Ident{name}
}

func (self *Ident) Type() string {
  return self.name
}

func (self *Ident) Eval(env *scope.Scope) value.Value {
  if obj := env.Lookup(self.name, true); obj != nil {
    if obj.Type == scope.Var {
      val := obj.Data.(value.Value)
      return val
    } else {
      return nil // TODO
    }
  }
  return nil
}

func (self *Ident) String() string {
  return self.name
}
