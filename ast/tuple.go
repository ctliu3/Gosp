package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
)

type Tuple struct {
  Nodes []Node
}

func NewTuple(nodes []Node) *Tuple {
  return &Tuple{Nodes: nodes}
}

func (self *Tuple) Type() string {
  return "tuple"
}

func (self *Tuple) Eval(env *scope.Scope) value.Value {
  return value.Int{Value: 5}
}

func (self *Tuple) String() string {
  return "tuple"
}
