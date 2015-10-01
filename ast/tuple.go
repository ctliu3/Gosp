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
  return nil
}

func (self *Tuple) String() string {
  var buf bytes.Buffer

  buf.WriteString("#<tuple>\n\t")
  for i, node := range self.Nodes {
    if i > 0 {
      buf.WriteString(", ")
    }
    buf.WriteString(node.Type())
  }
  return buf.String()
}

func (self *Tuple) ExtRep() string {
  var buf bytes.Buffer
  buf.WriteRune('(')
  for i, node := range self.Nodes {
    if i > 0 {
      buf.WriteRune(' ')
    }
    buf.WriteString(node.ExtRep())
  }
  buf.WriteRune(')')
  return buf.String()
}
