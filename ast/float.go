package ast

import (
  "fmt"
  "strconv"

  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
)

type Float struct {
  value float64
}

func NewFloat(name string) *Float {
  val, err := strconv.ParseFloat(name, 0)
  if err != nil {
    panic("float parse error")
  }
  return &Float{value: val}
}

func (self *Float) Type() string {
  return "float"
}

func (self *Float) Eval(env *scope.Scope) value.Value {
  return value.NewFloat(self.value)
}

func (self *Float) String() string {
  return fmt.Sprintf("%f", self.value)
}
