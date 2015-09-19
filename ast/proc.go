package ast

import (
  "fmt"
  "bytes"

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
    fmt.Printf("args, i %v type %v string %v\n", i, self.args[i].Type(), args[i].String())
    if args[i] == nil {
      panic(fmt.Errorf("parameter %v can not be found", self.args[i].(*Ident).Name))
    }
  }

  fmt.Printf("proc type %v\n", obj.Type)
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

  env := closure.Env.(*scope.Scope)
  formals := t.Formals.(*Tuple)
  body := t.Body.(Node)

  nargs := len(formals.Nodes)
  if nargs != len(args) {
    panic(fmt.Errorf(
      `the expected number of arguments does not match the given number
        expected: %v
        given: %v`, len(formals.Nodes), nargs))
  }

  for i := 0; i < nargs; i++ {
    env.Insert(formals.Nodes[i].(*Ident).Name, scope.NewObj(args[i]))
  }
  return body.Eval(env)
}

func (self *Proc) ProcName() string {
  return self.name
}

func (self *Proc) String() string {
  return fmt.Sprintf("#<procedure:%v>", self.name)
}

func (self *Proc) ExtRep() string {
  var buf bytes.Buffer

  buf.WriteString(fmt.Sprintf("(%v", self.name))
  for _, arg := range self.args {
    buf.WriteRune(' ')
    buf.WriteString(arg.ExtRep())
  }
  buf.WriteRune(')')

  return buf.String()
}
