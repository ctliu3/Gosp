package ast

import (
  "fmt"
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// (go (⟨operator⟩ ⟨operand1⟩ ...))

type Go struct {
  call Node
}

func NewGo(call Node) *Go {
  return &Go{call}
}

func (self *Go) Type() string {
  return const_.GO
}

func (self *Go) Eval(env *scope.Scope) value.Value {
  fmt.Println("Go#Eval")
  go self.call.Eval(env)
  return nil
}

func (self *Go) String() string {
  return "go"
}

func (self *Go) ExtRep() string {
  return ""
}
