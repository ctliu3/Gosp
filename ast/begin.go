package ast

import (
  "bytes"
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// (begin ⟨expression1 ⟩ ⟨expression2 ⟩ . . . )

type Begin struct {
  exprs []Node
}

func NewBegin(exprs []Node) *Begin {
  return &Begin{exprs}
}

func (self *Begin) Type() string {
  return const_.BEGIN
}

func (self *Begin) Eval(env *scope.Scope) value.Value {
  var ret value.Value
  for _, expr := range self.exprs {
    ret = expr.Eval(env)
  }
  return ret
}

func (self *Begin) String() string {
  return "begin"
}

func (self *Begin) ExtRep() string {
  var buffer bytes.Buffer

  buffer.WriteString("(begin")
  for _, node := range self.exprs {
    buffer.WriteString(node.ExtRep())
  }
  buffer.WriteString(")")

  return buffer.String()
}
