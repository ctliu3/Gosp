package ast

import (
  "fmt"
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

type Invoke struct {
  callee Node
  args []Node
}

func NewInvoke(callee Node, args []Node) *Invoke {
  return &Invoke{callee, args}
}

func (self *Invoke) Type() string {
  return const_.INVOKE
}

func (self *Invoke) Eval(env *scope.Scope) value.Value {
  proc := self.callee.(*Lambda).Eval(env)

  args := make([]value.Value, len(self.args))
  for i := 0; i < len(self.args); i++ {
    args[i] = self.args[i].Eval(env)
    if args[i] == nil {
      panic(fmt.Errorf("parameter %v can not be found", self.args[i].(*Ident).Name))
    }
  }
  return lambdaCall(proc.(*value.Closure), args...)
}

func (self *Invoke) String() string {
  return "invoke"
}

func (self *Invoke) ExtRep() string {
  return ""
}
