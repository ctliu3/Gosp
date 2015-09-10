package ast

import (
  "bytes"
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
  var buffer bytes.Buffer

  buffer.WriteString("#<tuple>\n\t")
  for i, node := range self.Nodes {
    if i > 0 {
      buffer.WriteString(", ")
    }
    buffer.WriteString(node.Type())
  }
  return buffer.String()
}
