package builtin

import (
  "fmt"
  "github.com/ctliu3/gosp/value"
)

// (display obj)

type Display struct {
  value.Proc
}

func NewDisplay() *Display {
  return &Display{}
}

func (self *Display) Call(args ...value.Value) value.Value {
  fmt.Printf(args[0].String())
  return nil
}

// (newline)

type Newline struct {
  value.Proc
}

func NewNewline() *Newline {
  return &Newline{}
}

func (self *Newline) Call(args ...value.Value) value.Value {
  fmt.Printf("\n")
  return nil
}
