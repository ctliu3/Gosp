package ast

import (
  "fmt"
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

type String struct {
  value string
}

func NewString(name string) *String {
  return &String{value: name}
}

func (self *String) Type() string {
  return const_.STRING
}

func (self *String) Eval(env *scope.Scope) value.Value {
  return value.NewString(self.value)
}

func (self *String) String() string {
  return self.value
}

func (self *String) ExtRep() string {
  return fmt.Sprintf("%v", self.value)
}
