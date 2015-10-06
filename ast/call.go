package ast

import (
  "fmt"
  "bytes"
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// (⟨operator⟩ ⟨operand1⟩ ...)

type Call struct {
  operator Node
  operands []Node
}

func NewCall(operator Node, operands []Node) *Call {
  return &Call{operator, operands}
}

func (self *Call) Type() string {
  return const_.CALL
}

func (self *Call) Eval(env *scope.Scope) value.Value {
  //fmt.Println("Call#Eval")

  obj := self.operator.Eval(env)

  args := make([]value.Value, len(self.operands))
  for i := 0; i < len(self.operands); i++ {
    args[i] = self.operands[i].Eval(env)
    if args[i] == nil {
      panic(fmt.Errorf("parameter %v can not be found", self.operands[i].(*Ident).Name))
    }
  }

  switch obj.(type) {
  case value.Builtin:
    return obj.(value.Builtin).Call(args...)
  case *value.Closure:
    return precedureCall(obj.(*value.Closure), args...)
  default:
    panic("unexpected procedure")
  }

  return nil
}

func precedureCall(closure *value.Closure, args ...value.Value) value.Value {
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

func (self *Call) String() string {
  return "precedure call"
}

func (self *Call) ExtRep() string {
  var buf bytes.Buffer

  buf.WriteRune('(')
  buf.WriteString(self.operator.ExtRep())
  for _, arg := range self.operands {
    buf.WriteRune(' ')
    buf.WriteString(arg.ExtRep())
  }
  buf.WriteRune(')')

  return buf.String()
}
