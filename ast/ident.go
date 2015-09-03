package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
)

type Ident struct {
  Name string
}

func NewIdent(name string) *Ident {
  return &Ident{Name: name}
}

func (self *Ident) Type() string {
  return self.Name
}

func (self *Ident) Eval(env *scope.Scope) value.Value {
  return value.NewInt(5)
  //var val value.Int
  //if ok := env.Lookup(self.name); ok {
    //return true
  //}
  //return false
}

func (self *Ident) String() string {
  return ""
}
