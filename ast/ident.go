package ast

import (
  "fmt"
  //"reflect"

  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
)

type Ident struct {
  Name string
}

func NewIdent(name string) *Ident {
  return &Ident{name}
}

func (self *Ident) Type() string {
  return self.Name
}

func (self *Ident) Eval(env *scope.Scope) value.Value {
  if obj := env.Lookup(self.Name, true); obj != nil {
    //fmt.Println(reflect.TypeOf(obj.Data))
    return obj.Data.(value.Value)
  }
  return nil
}

func (self *Ident) String() string {
  return fmt.Sprintf("#<procedure:%v>", self.Name)
}

func (self *Ident) ExtRep() string {
  return self.Name
}
