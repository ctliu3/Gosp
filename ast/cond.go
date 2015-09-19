package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// (cond ⟨clause1⟩ ⟨clause2⟩ . . . )
// where each ⟨clause⟩ should be (⟨test⟩ ⟨expression1⟩ ...)

type Cond struct {
  clause []Tuple
}

func NewCond(clause []Tuple) *Cond {
  return &Cond{clause}
}

func (self *Cond) Type() string {
  return const_.COND
}

func (self *Cond) Eval(env *scope.Scope) value.Value {
  var ret value.Value

  nclause := len(self.clause)
  for i, node := range self.clause {
    if i == nclause - 1 {
      if node.Nodes[0].Type() == const_.ELSE {
        return node.Nodes[1].Eval(env)
      }
    }
    ret = node.Nodes[0].Eval(env)
    if ret.String() != const_.TRUE && ret.String() != const_.FALSE {
      panic("cond: bad syntax?")
    }
    if ret.String() == const_.TRUE {
      return node.Nodes[1].Eval(env)
    }
  }
  return nil
}

func (self *Cond) String() string {
  return "cond"
}

func (self *Cond) ExtRep() string {
  return ""
}
