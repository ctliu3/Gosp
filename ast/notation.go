package ast

import (
  "fmt"

  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

type Notation struct {
  name string
}

func NewNotation(name string) *Notation {
  return &Notation{name}
}

func (self *Notation) Type() string {
  return const_.NOTATION
}

func (self *Notation) Eval(env *scope.Scope) value.Value {
  panic(fmt.Errorf("%v: can not be evaluated.", self.name))
  return nil
}

func (self *Notation) String() string {
  return self.name
}

func (self *Notation) ExtRep() string {
  return self.String()
}
