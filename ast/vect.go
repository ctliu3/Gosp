package ast

import (
  //"fmt"
  "bytes"
  //"reflect"

  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

type Vect struct {
  Nodes []Node
}

func NewVect(nodes []Node) *Vect {
  return &Vect{Nodes: nodes}
}

func (self *Vect) Type() string {
  return const_.VECT
}

func (self *Vect) Eval(env *scope.Scope) value.Value {
  elements := make([]value.Value, len(self.Nodes))
  for i, node := range self.Nodes {
    elements[i] = node.Eval(env)
  }
  return value.NewVect(elements)
}

func (self *Vect) String() string {
  var buf bytes.Buffer

  buf.WriteString("#<vector>\n\t")
  for i, node := range self.Nodes {
    if i > 0 {
      buf.WriteString(", ")
    }
    buf.WriteString(node.Type())
  }

  return buf.String()
}

func (self *Vect) ExtRep() string {
  var buf bytes.Buffer

  buf.WriteRune('[')
  for i, node := range self.Nodes {
    if i > 0 {
      buf.WriteRune(' ')
    }
    buf.WriteString(node.ExtRep())
  }
  buf.WriteRune(']')

  return buf.String()
}
