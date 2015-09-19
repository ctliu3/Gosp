package ast

import (
  "github.com/ctliu3/gosp/scope"
  "github.com/ctliu3/gosp/value"
  const_ "github.com/ctliu3/gosp/constant"
)

// (case ⟨key⟩ ⟨clause1 ⟩ ⟨clause2 ⟩ . . . )

type Case struct {
  key Node
  clause []Tuple
}

func NewCase(key Node, clause []Tuple) *Case {
  return &Case{key, clause}
}

func (self *Case) Type() string {
  return const_.CASE
}

func (self *Case) Eval(env *scope.Scope) value.Value {
  //var ret value.Value

  //nclause := len(self.clause)
  //for i, node := range self.clause {
    //if i == nclause - 1 {
      //if node.Nodes[0].Type() == const_.ELSE {
        //return node.Nodes[1].Eval(env)
      //}
    //}
    //ret = node.Nodes[0].Eval(env)
    //if ret.String() != const_.TRUE && ret.String() != const_.FALSE {
      //panic("cond: bad syntax?")
    //}
    //if ret.String() == const_.TRUE {
      //return node.Nodes[1].Eval(env)
    //}
  //}
  return nil
}

func (self *Case) String() string {
  return "case"
}

func (self *Case) ExtRep() string {
  return ""
}
