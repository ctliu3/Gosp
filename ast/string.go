package ast

import (
  "fmt"
  "strconv"

  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
)

type String struct {
  value string
}

func NewString(name string) *String {
  return &String{value: name}
}

func (self *String) Type() string {
  return "string"
}

func (self *String) Eval(env *scope.Scope) value.Value {
  return value.NewString(self.value)
}

func (self *String) String() string {
  return fmt.Sprintf("%d", self.value)
}
