package ast

import (
  "fmt"
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// (display obj) 

type Display struct {
  obj Node
}

func NewDisplay(obj Node) *Display {
  return &Display{obj}
}

func (self *Display) Type() string {
  return const_.DISPLAY
}

func (self *Display) Eval(env *scope.Scope) value.Value {
  fmt.Println("Display#Eval")
  fmt.Printf("%v", self.obj.Eval(env).String())
  return nil
}

func (self *Display) ExtRep() string {
  return ""
}
