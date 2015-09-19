package ast

import (
  "fmt"
  "strconv"

  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

type Float struct {
  value float64
}

func NewFloat(name string) *Float {
  val, err := strconv.ParseFloat(name, 0)
  if err != nil {
    panic(err)
  }
  return &Float{value: val}
}

func (self *Float) Type() string {
  return const_.FLOAT
}

func (self *Float) Eval(env *scope.Scope) value.Value {
  return value.NewFloat(self.value)
}

func (self *Float) String() string {
  return fmt.Sprintf("%v", self.value)
}

func (self *Float) ExtRep() string {
  return ""
}
