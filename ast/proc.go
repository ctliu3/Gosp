package ast

import (
  "fmt"

  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  "github.com/ctliu3/gosp/procs"
  const_ "github.com/ctliu3/gosp/constant"
)

type Proc struct {
  name string
  args []Node
}

func NewProc(name string, args []Node) *Proc {
  return &Proc{name: name, args: args}
}

func (self *Proc) Type() string {
  return const_.PROC
}

func (self *Proc) Eval(env *scope.Scope) value.Value {
  obj := env.Lookup(self.name, true)

  if obj == nil {
    panic("undefined procedure")
  }
  switch obj.Type {
  case scope.Proc:
    proc := obj.Data.(procs.Proc)
    args := make([]value.Value, len(self.args))
    for i := 0; i < len(self.args); i++ {
      args[i] = self.args[i].Eval(env)
      fmt.Println("#" + args[i].String())
    }

    return proc.Call(args...)
  case scope.Var:
    panic(self.name + " should be procedure")
  }
  return nil
}

func (self *Proc) String() string {
  return self.name
}
