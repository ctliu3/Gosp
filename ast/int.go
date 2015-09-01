package ast

import (
  "fmt"
  "strconv"

  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
)

type Int struct {
  value int64
}

func NewInt(name string) *Int {
  val, err := strconv.ParseInt(name, 10, 0)
  if err != nil {
    panic("int parse error")
  }
  return &Int{value: val}
}

func (self *Int) Type() string {
  return "integer"
}

func (self *Int) Eval(env *scope.Scope) value.Value {
  return value.Int{Value: self.value}
}

func (self *Int) String() string {
  return fmt.Sprintf("%d", self.value)
}
