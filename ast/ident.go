package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
)

type Ident struct {
  name string
}

func NewIdent(name string) *Ident {
  return &Ident{name: name}
}

func (self *Ident) Type() string {
  return self.name
}

func (self *Ident) Eval(env *scope.Scope) value.Value {
  return value.Int{Value: 5}
  //var val value.Int
  //if ok := env.Lookup(self.name); ok {
    //return true
  //}
  //return false
}

func (self *Ident) String() string {
  return ""
}
