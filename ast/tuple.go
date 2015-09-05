package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// If one Node is Tuple, it means that it is surrounded by `()'.
type Tuple struct {
  Nodes []Node
}

func NewTuple(nodes []Node) *Tuple {
  return &Tuple{Nodes: nodes}
}

func (self *Tuple) Type() string {
  return const_.TUPLE
}

func (self *Tuple) Eval(env *scope.Scope) value.Value {
  return value.NewInt(5)
}

func (self *Tuple) String() string {
  return const_.TUPLE
}
