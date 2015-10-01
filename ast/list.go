package ast

import (
  //"fmt"
  "bytes"
  //"reflect"

  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

type List struct {
  Nodes []Node
}

func NewList(nodes []Node) *List {
  return &List{Nodes: nodes}
}

func (self *List) Type() string {
  return const_.LIST
}

func (self *List) Eval(env *scope.Scope) value.Value {
  elements := make([]value.Value, len(self.Nodes))
  for i, node := range self.Nodes {
    elements[i] = node.Eval(env)
  }
  return value.NewList(elements)
}

func (self *List) String() string {
  var buf bytes.Buffer

  buf.WriteString("#<list>\n\t")
  for i, node := range self.Nodes {
    if i > 0 {
      buf.WriteString(", ")
    }
    buf.WriteString(node.Type())
  }
  return buf.String()
}

func (self *List) ExtRep() string {
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
