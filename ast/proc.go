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
  fmt.Println("Proc#Eval")
  obj := env.Lookup(self.name, true)

  if obj == nil {
    panic(fmt.Errorf(
      "%v: undefined;\n cannot reference undefined identifier", self.name))
  }

  args := make([]value.Value, len(self.args))
  for i := 0; i < len(self.args); i++ {
    args[i] = self.args[i].Eval(env)
  }
  switch obj.Type {
  case scope.Proc:
    proc := obj.Data.(procs.Proc)
    return proc.Call(args...)
  case scope.Var:
    switch obj.Data.(type) {
    case *value.Closure:
      return lambdaCall(obj.Data.(*value.Closure), args...)
    default:
      panic(fmt.Errorf("%v should be identifier", self.name))
    }
  }
  return nil
}

func lambdaCall(closure *value.Closure, args ...value.Value) value.Value {
  t := closure.Lambda.(*Lambda)
  formals := t.Formals.(*Tuple)
  body := t.Body.(Node)

  nargs := len(formals.Nodes)
  if nargs != len(args) {
    panic(fmt.Errorf(
      `the expected number of arguments does not match the given number
        expected: %v
        given: %v`, len(formals.Nodes), nargs))
  }

  env := scope.NewScope(nil)
  for i := 0; i < nargs; i++ {
    env.Insert(formals.Nodes[i].(*Ident).Name, scope.NewObj(args[i]))
  }
  return body.Eval(env)
}

func (self *Proc) String() string {
  return fmt.Sprintf("#<procedure:%v>", self.name)
}
