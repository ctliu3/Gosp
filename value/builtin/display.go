package builtin

import (
  "fmt"
  "github.com/ctliu3/gosp/value"
)

type Display struct {
  value.Proc
}

func NewDisplay() *Display {
  return &Display{}
}

func (self *Display) Call(args ...value.Value) value.Value {
  fmt.Println(args[0].String())
  return nil
}
